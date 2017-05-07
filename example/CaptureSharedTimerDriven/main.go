// +build windows
package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca"
)

type WAVEFormat struct {
	FormatTag      uint16
	Channels       uint16
	SamplesPerSec  uint32
	AvgBytesPerSec uint32
	BlockAlign     uint16
	BitsPerSample  uint16
	DataTag        [4]byte // 'data'
	DataSize       uint32
	RawData        []byte
}

func (v *WAVEFormat) Bytes() (output []byte) {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.BigEndian, []byte("RIFF"))
	binary.Write(buf, binary.LittleEndian, uint32(v.DataSize+36)) // Header size is 44 byte, so 44 - 8 = 36
	binary.Write(buf, binary.BigEndian, []byte("WAVEfmt "))
	binary.Write(buf, binary.LittleEndian, uint32(16)) // 16 (0x10000000) for PCM
	binary.Write(buf, binary.LittleEndian, uint16(1))  // 1 (0x0001) for PCM
	binary.Write(buf, binary.LittleEndian, v.Channels)
	binary.Write(buf, binary.LittleEndian, v.SamplesPerSec)
	binary.Write(buf, binary.LittleEndian, v.AvgBytesPerSec)
	binary.Write(buf, binary.LittleEndian, v.BlockAlign)
	binary.Write(buf, binary.LittleEndian, v.BitsPerSample)
	binary.Write(buf, binary.BigEndian, []byte("data"))
	binary.Write(buf, binary.LittleEndian, v.DataSize)
	binary.Write(buf, binary.LittleEndian, v.RawData)

	return buf.Bytes()
}

func main() {
	var err error
	if err = run(os.Args); err != nil {
		log.Fatal(err)
	}
	return
}

func run(args []string) (err error) {
	var durationFlag int64
	var filenameFlag string
	var audio *WAVEFormat

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Int64Var(&durationFlag, "t", 0, "Specify recording time in millisecond")
	f.StringVar(&filenameFlag, "f", "output.wav", "Specify file name to save (default is output.wav)")
	f.Parse(args[1:])

	if durationFlag <= 0 {
		return
	}
	if audio, err = captureSharedTimerDriven(); err != nil {
		return
	}
	if err = ioutil.WriteFile(filenameFlag, audio.Bytes(), 0644); err != nil {
		return
	}
	fmt.Printf("Saved captured audio as %s\n", filenameFlag)
	return
}

func captureSharedTimerDriven() (audio *WAVEFormat, err error) {
	audio = &WAVEFormat{}

	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}

	var mmde *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err = mmde.GetDefaultAudioEndpoint(wca.ECapture, wca.EConsole, &mmd); err != nil {
		return
	}
	defer mmd.Release()

	var ps *wca.IPropertyStore
	if err = mmd.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		return
	}
	defer ps.Release()

	var pv wca.PROPVARIANT
	if err = ps.GetValue(&wca.PKEY_Device_FriendlyName, &pv); err != nil {
		return
	}
	fmt.Printf("capturing from\n%s\n", pv.String())

	var ac *wca.IAudioClient
	if err = mmd.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &ac); err != nil {
		return
	}
	defer ac.Release()

	var wfx *wca.WAVEFORMATEX
	if err = ac.GetMixFormat(&wfx); err != nil {
		return
	}
	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(wfx)))

	wfx.WFormatTag = 1
	wfx.WBitsPerSample = 16
	wfx.NSamplesPerSec = 48000
	wfx.NBlockAlign = (wfx.WBitsPerSample / 8) * wfx.NChannels
	wfx.NAvgBytesPerSec = wfx.NSamplesPerSec * uint32(wfx.NBlockAlign)
	wfx.CbSize = 0

	audio.Channels = wfx.NChannels
	audio.SamplesPerSec = wfx.NSamplesPerSec
	audio.AvgBytesPerSec = wfx.NAvgBytesPerSec
	audio.BlockAlign = wfx.NBlockAlign
	audio.BitsPerSample = wfx.WBitsPerSample

	fmt.Println("--------")
	fmt.Printf("Format: PCM %d bit signed integer\n", wfx.WBitsPerSample)
	fmt.Printf("Rate: %d Hz\n", wfx.NSamplesPerSec)
	fmt.Printf("Channels: %d\n", wfx.NChannels)
	fmt.Println("--------")

	var defaultPeriod int64
	var minimumPeriod int64
	var capturingPeriod time.Duration
	if err = ac.GetDevicePeriod(&defaultPeriod, &minimumPeriod); err != nil {
		return
	}
	capturingPeriod = time.Duration(int(defaultPeriod) * 100)
	fmt.Printf("Default capturing period: %d ms\n", capturingPeriod/time.Millisecond)

	if err = ac.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, 0, 250*10000, 0, wfx, nil); err != nil {
		return
	}

	var bufferFrameSize uint32
	if err = ac.GetBufferSize(&bufferFrameSize); err != nil {
		return
	}
	fmt.Printf("Allocated buffer size: %d\n", bufferFrameSize)

	var acc *wca.IAudioCaptureClient
	if err = ac.GetService(wca.IID_IAudioCaptureClient, &acc); err != nil {
		return
	}
	defer acc.Release()

	if err = ac.Start(); err != nil {
		return
	}
	fmt.Println("Start capturing loopback audio")

	time.Sleep(capturingPeriod)

	var isCapturing bool
	var data *byte
	var availableFrameSize uint32
	var flags uint32
	var devicePosition uint64
	var qcpPosition uint64
	var b *byte
	var padding uint32

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	isCapturing = true

	for {
		if !isCapturing {
			break
		}
		select {
		case <-signalChan:
			fmt.Println("interrupted by signal")
			isCapturing = false
			break
		default:
			if err = acc.GetBuffer(&data, &availableFrameSize, &flags, &devicePosition, &qcpPosition); err != nil {
				return
			}
			if availableFrameSize == 0 {
				continue
			}

			start := unsafe.Pointer(data)
			lim := int(availableFrameSize) * int(wfx.NBlockAlign)
			for n := 0; n < lim; n++ {
				b = (*byte)(unsafe.Pointer(uintptr(start) + uintptr(n)))
				audio.RawData = append(audio.RawData, *b)
			}
			audio.DataSize += uint32(lim)
			if err = ac.GetCurrentPadding(&padding); err != nil {
				return
			}
			//capturingPeriod = time.Duration(1000000 * 1000 * int(bufferFrameSize-padding) / int(wfx.NSamplesPerSec))
			//time.Sleep(capturingPeriod / 2)
			time.Sleep(capturingPeriod)
			if err = acc.ReleaseBuffer(availableFrameSize); err != nil {
				return
			}
		}
	}
	//audio.RawData = output

	if err = ac.Stop(); err != nil {
		return
	}
	fmt.Println("Stopping capturing loopback audio")
	return
}
