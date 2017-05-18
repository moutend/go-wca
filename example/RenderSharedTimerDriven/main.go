// +build windows
package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
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
	DataSize       uint32
	RawData        []byte
}

type FilenameFlag struct {
	Value string
}

func (f *FilenameFlag) Set(value string) (err error) {
	if !strings.HasSuffix(value, ".wav") {
		err = fmt.Errorf("specify WAVE audio file (*.wav)")
		return
	}
	f.Value = value
	return
}

func (f *FilenameFlag) String() string {
	return f.Value
}

func readFile(filename string) (audio *WAVEFormat, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	audio = &WAVEFormat{}
	reader := bytes.NewReader(file)
	binary.Read(io.NewSectionReader(reader, 20, 2), binary.LittleEndian, &audio.FormatTag)
	binary.Read(io.NewSectionReader(reader, 22, 2), binary.LittleEndian, &audio.Channels)
	binary.Read(io.NewSectionReader(reader, 24, 4), binary.LittleEndian, &audio.SamplesPerSec)
	binary.Read(io.NewSectionReader(reader, 28, 4), binary.LittleEndian, &audio.AvgBytesPerSec)
	binary.Read(io.NewSectionReader(reader, 32, 2), binary.LittleEndian, &audio.BlockAlign)
	binary.Read(io.NewSectionReader(reader, 34, 2), binary.LittleEndian, &audio.BitsPerSample)
	binary.Read(io.NewSectionReader(reader, 40, 4), binary.LittleEndian, &audio.DataSize)

	buf := new(bytes.Buffer)
	io.Copy(buf, io.NewSectionReader(reader, 44, int64(audio.DataSize)))
	audio.RawData = buf.Bytes()

	if len(audio.RawData) == 0 {
		err = fmt.Errorf("empty data")
	}
	return
}

func main() {
	var err error
	if err = run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) (err error) {
	var filenameFlag FilenameFlag
	var audio *WAVEFormat

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Var(&filenameFlag, "input", "Specify WAVE format audio (e.g. music.wav)")
	f.Var(&filenameFlag, "i", "Alias of --input")
	f.Parse(args[1:])

	if filenameFlag.Value == "" {
		return
	}
	if audio, err = readFile(filenameFlag.Value); err != nil {
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-signalChan:
			fmt.Println("Interrupted by SIGINT")
			cancel()
		}
		return
	}()

	if err = renderSharedTimerDriven(ctx, audio); err != nil {
		return
	}
	fmt.Println("Successfully done")
	return
}

func renderSharedTimerDriven(ctx context.Context, audio *WAVEFormat) (err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}
	defer ole.CoUninitialize()

	var de *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &de); err != nil {
		return
	}
	defer de.Release()

	var mmd *wca.IMMDevice
	if err = de.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
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
	fmt.Printf("Rendering audio to: %s\n", pv.String())

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

	if wfx.WFormatTag != wca.WAVE_FORMAT_PCM {
		wfx.WFormatTag = 1
		wfx.NSamplesPerSec = audio.SamplesPerSec
		wfx.WBitsPerSample = audio.BitsPerSample
		wfx.NChannels = audio.Channels
		wfx.NBlockAlign = audio.BlockAlign
		wfx.NAvgBytesPerSec = audio.AvgBytesPerSec
		wfx.CbSize = 0
	}

	fmt.Println("--------")
	fmt.Printf("Format: PCM %d bit signed integer\n", wfx.WBitsPerSample)
	fmt.Printf("Rate: %d Hz\n", wfx.NSamplesPerSec)
	fmt.Printf("Channels: %d\n", wfx.NChannels)
	fmt.Println("--------")

	var defaultPeriod int64
	var minimumPeriod int64
	var renderingPeriod time.Duration
	if err = ac.GetDevicePeriod(&defaultPeriod, &minimumPeriod); err != nil {
		return
	}
	renderingPeriod = time.Duration(int(defaultPeriod) * 100)
	fmt.Printf("Default rendering period: %d ms\n", renderingPeriod/time.Millisecond)

	if err = ac.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, 0, 200*10000, 0, wfx, nil); err != nil {
		return
	}

	var bufferFrameSize uint32
	if err = ac.GetBufferSize(&bufferFrameSize); err != nil {
		return
	}
	fmt.Printf("Allocated buffer size: %d\n", bufferFrameSize)

	var arc *wca.IAudioRenderClient
	if err = ac.GetService(wca.IID_IAudioRenderClient, &arc); err != nil {
		return
	}
	defer arc.Release()

	if err = ac.Start(); err != nil {
		return
	}
	fmt.Println("Start rendering audio with shared-timer-driven mode")
	fmt.Println("Press Ctrl-C to stop rendering")

	time.Sleep(renderingPeriod)

	var data *byte
	var offset int
	var padding uint32
	var availableFrameSize uint32
	var b *byte
	var isPlaying bool = true

	for {
		if !isPlaying {
			break
		}
		select {
		case <-ctx.Done():
			isPlaying = false
			break
		default:
			if offset >= int(audio.DataSize) {
				isPlaying = false
				break
			}
			if err = ac.GetCurrentPadding(&padding); err != nil {
				return
			}
			availableFrameSize = bufferFrameSize - padding
			if availableFrameSize == 0 {
				continue
			}
			if err = arc.GetBuffer(availableFrameSize, &data); err != nil {
				return
			}

			start := unsafe.Pointer(data)
			lim := int(availableFrameSize) * int(wfx.NBlockAlign)
			remaining := int(audio.DataSize) - offset
			if remaining < lim {
				lim = remaining
			}
			for n := 0; n < lim; n++ {
				b = (*byte)(unsafe.Pointer(uintptr(start) + uintptr(n)))
				*b = audio.RawData[offset+n]
			}
			offset += lim
			if err = arc.ReleaseBuffer(availableFrameSize, 0); err != nil {
				return
			}
			time.Sleep(renderingPeriod)
		}
	}

	if err = ac.Stop(); err != nil {
		return
	}
	fmt.Println("Stop rendering loopback audio")
	return
}
