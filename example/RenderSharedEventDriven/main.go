// +build windows
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wave"
	"github.com/moutend/go-wca"
)

var version = "latest"
var revision = "latest"

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
	var filenameFlag FilenameFlag
	var versionFlag bool
	var audio *wave.WAVE

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Var(&filenameFlag, "input", "Specify WAVE format audio (e.g. music.wav)")
	f.Var(&filenameFlag, "i", "Alias of --input")
	f.BoolVar(&versionFlag, "version", false, "Show version")
	f.Parse(args[1:])

	if versionFlag {
		fmt.Printf("%s-%s\n", version, revision)
		return
	}
	if filenameFlag.Value == "" {
		return
	}
	if audio, err = wave.OpenFile(filenameFlag.Value); err != nil {
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
	}()

	if err = renderSharedEventDriven(ctx, audio); err != nil {
		return
	}
	fmt.Println("Successfully done")
	return
}

func renderSharedEventDriven(ctx context.Context, audio *wave.WAVE) (err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}
	defer ole.CoUninitialize()

	var mmde *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err = mmde.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
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
	renderingPeriod = time.Duration(defaultPeriod * 100)
	fmt.Printf("Default rendering period: %d ms\n", renderingPeriod/time.Millisecond)

	if err = ac.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, wca.AUDCLNT_STREAMFLAGS_EVENTCALLBACK, 200*10000, 0, wfx, nil); err != nil {
		return
	}
	audioReadyEvent := wca.CreateEventExA(0, 0, 0, wca.EVENT_MODIFY_STATE|wca.SYNCHRONIZE)
	defer wca.CloseHandle(audioReadyEvent)

	if err = ac.SetEventHandle(audioReadyEvent); err != nil {
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
	fmt.Println("Start rendering audio with shared-event-driven mode")
	fmt.Println("Press Ctrl-C to stop")
	var b *byte
	var data *byte
	var offset int
	var isPlaying bool = true
	var padding uint32
	var availableFrameSize uint32

	errorChan := make(chan error, 1)

	for {
		if !isPlaying {
			close(errorChan)
			break
		}
		go func() {
			errorChan <- watchEvent(ctx, audioReadyEvent)
		}()
		select {
		case <-ctx.Done():
			isPlaying = false
			<-errorChan
			break
		case err = <-errorChan:
			if err != nil {
				isPlaying = false
				break
			}
			if offset >= int(audio.DataSize) {
				isPlaying = false
				break
			}
			if err = ac.GetCurrentPadding(&padding); err != nil {
				return
			}
			if availableFrameSize = bufferFrameSize - padding; availableFrameSize == 0 {
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
		}
	}
	if err != nil {
		return
	}

	if err = ac.Stop(); err != nil {
		return
	}
	fmt.Println("Stop rendering")
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
	//if err = ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
	//	return
	//}
	dw := wca.WaitForSingleObject(event, wca.INFINITE)
	if dw != 0 {
		return fmt.Errorf("failed to watch event")
	}
	//ole.CoUninitialize()
	return
}
