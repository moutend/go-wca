// +build windows
package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
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

type DurationFlag struct {
	Value time.Duration
}

func (f *DurationFlag) Set(value string) (err error) {
	var sec float64

	if sec, err = strconv.ParseFloat(value, 64); err != nil {
		return
	}
	f.Value = time.Duration(sec * float64(time.Second))
	return
}

func (f *DurationFlag) String() string {
	return f.Value.String()
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

func main() {
	var err error
	if err = run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) (err error) {
	var durationFlag DurationFlag
	var filenameFlag FilenameFlag
	var audio *WAVEFormat

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Var(&durationFlag, "duration", "Specify recording duration in second")
	f.Var(&durationFlag, "d", "Alias of --duration")
	f.Var(&filenameFlag, "output", "file name")
	f.Var(&filenameFlag, "o", "Alias of --output")
	f.Parse(args[1:])

	if filenameFlag.Value == "" {
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

	if audio, err = loopbackCaptureSharedEventDriven(ctx, durationFlag.Value); err != nil {
		return
	}
	if err = ioutil.WriteFile(filenameFlag.Value, audio.Bytes(), 0644); err != nil {
		return
	}
	fmt.Println("Successfully done")
	return
}

func loopbackCaptureSharedEventDriven(ctx context.Context, duration time.Duration) (audio *WAVEFormat, err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}

	var mmdCapturee *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmdCapturee); err != nil {
		return
	}
	defer mmdCapturee.Release()

	var mmdCapture *wca.IMMDevice
	if err = mmdCapturee.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmdCapture); err != nil {
		return
	}
	defer mmdCapture.Release()

	var mmdRender *wca.IMMDevice
	if err = mmdCapturee.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmdRender); err != nil {
		return
	}
	defer mmdRender.Release()

	var ps *wca.IPropertyStore
	if err = mmdCapture.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		return
	}
	defer ps.Release()

	var pv wca.PROPVARIANT
	if err = ps.GetValue(&wca.PKEY_Device_FriendlyName, &pv); err != nil {
		return
	}
	fmt.Printf("Capturing audio from: %s\n", pv.String())

	var acCapture *wca.IAudioClient
	if err = mmdCapture.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &acCapture); err != nil {
		return
	}
	defer acCapture.Release()

	var acRender *wca.IAudioClient
	if err = mmdRender.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &acRender); err != nil {
		return
	}
	defer acRender.Release()

	var wfx *wca.WAVEFORMATEX
	if err = acCapture.GetMixFormat(&wfx); err != nil {
		return
	}
	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(wfx)))

	wfx.WFormatTag = 1
	wfx.NChannels = 2
	wfx.NSamplesPerSec = 44100
	wfx.WBitsPerSample = 16
	wfx.NBlockAlign = (wfx.WBitsPerSample / 8) * wfx.NChannels
	wfx.NAvgBytesPerSec = wfx.NSamplesPerSec * uint32(wfx.NBlockAlign)
	wfx.CbSize = 0

	audio = &WAVEFormat{}
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
	if err = acCapture.GetDevicePeriod(&defaultPeriod, &minimumPeriod); err != nil {
		return
	}
	capturingPeriod = time.Duration(int(defaultPeriod) * 100)
	fmt.Printf("Default capturing period: %d ms\n", capturingPeriod/time.Millisecond)

	if err = acCapture.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, wca.AUDCLNT_STREAMFLAGS_EVENTCALLBACK|wca.AUDCLNT_STREAMFLAGS_LOOPBACK, 200*10000, 0, wfx, nil); err != nil {
		return
	}
	if err = acRender.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, wca.AUDCLNT_STREAMFLAGS_EVENTCALLBACK, 200*10000, 0, wfx, nil); err != nil {
		return
	}

	audioReadyEvent := wca.CreateEventExA(0, 0, 0, wca.EVENT_MODIFY_STATE|wca.SYNCHRONIZE)
	defer wca.CloseHandle(audioReadyEvent)

	fakeAudioReadyEvent := wca.CreateEventExA(0, 0, 0, wca.EVENT_MODIFY_STATE|wca.SYNCHRONIZE)
	defer wca.CloseHandle(fakeAudioReadyEvent)
	if err = acCapture.SetEventHandle(fakeAudioReadyEvent); err != nil {
		return
	}
	if err = acRender.SetEventHandle(audioReadyEvent); err != nil {
		return
	}
	var bufferFrameSizeRender uint32
	if err = acRender.GetBufferSize(&bufferFrameSizeRender); err != nil {
		return
	}

	var bufferFrameSize uint32
	if err = acCapture.GetBufferSize(&bufferFrameSize); err != nil {
		return
	}

	fmt.Printf("Allocated buffer size: %d\n", bufferFrameSize)

	var arc *wca.IAudioRenderClient
	if err = acRender.GetService(wca.IID_IAudioRenderClient, &arc); err != nil {
		return
	}
	defer arc.Release()

	var acc *wca.IAudioCaptureClient
	if err = acCapture.GetService(wca.IID_IAudioCaptureClient, &acc); err != nil {
		return
	}
	defer acc.Release()

	if err = acRender.Start(); err != nil {
		return
	}
	if err = acCapture.Start(); err != nil {
		return
	}

	fmt.Println("Start capturing loopback audio with shared-event-driven mode")

	if duration <= 0 {
		fmt.Println("Press Ctrl-C to stop capturing")
	}

	var isCapturing bool = true
	var currentDuration time.Duration
	var availableFrameSizeRender uint32
	var paddingRender uint32
	var dataRender *byte
	var data *byte
	var b *byte
	var availableFrameSize uint32
	var flags uint32
	var devicePosition uint64
	var qcpPosition uint64

	errorChan := make(chan error, 1)
	// Render silence
	if err = acRender.GetCurrentPadding(&paddingRender); err != nil {
		return
	}
	if availableFrameSizeRender = bufferFrameSizeRender - paddingRender; availableFrameSizeRender == 0 {
		fmt.Println("oops")
	}
	if err = arc.GetBuffer(availableFrameSizeRender, &dataRender); err != nil {
		return
	}

	startRender := unsafe.Pointer(dataRender)
	limRender := int(availableFrameSize) * int(wfx.NBlockAlign)

	for n := 0; n < limRender; n++ {
		b = (*byte)(unsafe.Pointer(uintptr(startRender) + uintptr(n)))
		*b = 0 // 0 indicates silence
	}
	if err = arc.ReleaseBuffer(availableFrameSizeRender, 0); err != nil {
		return
	}

	for {
		if !isCapturing {
			close(errorChan)
			break
		}
		go func() {
			errorChan <- watchEvent(ctx, audioReadyEvent)
		}()
		select {
		case <-ctx.Done():
			isCapturing = false
			<-errorChan
			break
		case err = <-errorChan:
			currentDuration = time.Duration(float64(audio.DataSize) / float64(audio.BitsPerSample/8) / float64(audio.Channels) / float64(audio.SamplesPerSec) * float64(time.Second))
			if duration != 0 && currentDuration > duration {
				isCapturing = false
				break
			}
			if err != nil {
				isCapturing = false
				break
			}

			// Render silence
			if err = acRender.GetCurrentPadding(&paddingRender); err != nil {
				return
			}
			if availableFrameSizeRender = bufferFrameSizeRender - paddingRender; availableFrameSizeRender == 0 {
				continue
			}
			if err = arc.GetBuffer(availableFrameSizeRender, &dataRender); err != nil {
				return
			}

			startRender := unsafe.Pointer(dataRender)
			limRender := int(availableFrameSize) * int(wfx.NBlockAlign)

			for n := 0; n < limRender; n++ {
				b = (*byte)(unsafe.Pointer(uintptr(startRender) + uintptr(n)))
				*b = 0 // 0 indicates silence
			}

			if err = arc.ReleaseBuffer(availableFrameSizeRender, 0); err != nil {
				return
			}

			// Capture loopback
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

			if err = acc.ReleaseBuffer(availableFrameSize); err != nil {
				return
			}
		}
	}

	fmt.Println("Stop capturing")
	if err = acCapture.Stop(); err != nil {
		return
	}
	if err = acRender.Stop(); err != nil {
		return
	}
	return
}

func watchEvent(ctx context.Context, event uintptr) (err error) {
	errorChan := make(chan error, 1)
	go func() {
		errorChan <- eventEmitter(event)
	}()
	select {
	case err = <-errorChan:
		close(errorChan)
		return
	case <-ctx.Done():
		err = ctx.Err()
		return
	}
	return
}

func eventEmitter(event uintptr) (err error) {
	dw := wca.WaitForSingleObject(event, wca.INFINITE)
	if dw != 0 {
		return fmt.Errorf("failed to watch event")
	}
	return
}
