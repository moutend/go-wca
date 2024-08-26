package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wav"
	"github.com/moutend/go-wca/pkg/wca"
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
	var audio = &wav.File{}
	var file []byte

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
	if file, err = ioutil.ReadFile(filenameFlag.Value); err != nil {
		return
	}
	if err = wav.Unmarshal(file, audio); err != nil {
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

func renderSharedTimerDriven(ctx context.Context, audio *wav.File) (err error) {
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

	var ac3 *wca.IAudioClient3
	if err = mmd.Activate(wca.IID_IAudioClient3, wca.CLSCTX_ALL, nil, &ac3); err != nil {
		return
	}
	defer ac3.Release()

	var wfx *wca.WAVEFORMATEX
	if err = ac3.GetMixFormat(&wfx); err != nil {
		return
	}
	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(wfx)))

	wfx.WFormatTag = 1
	wfx.NSamplesPerSec = uint32(audio.SamplesPerSec())
	wfx.WBitsPerSample = uint16(audio.BitsPerSample())
	wfx.NChannels = uint16(audio.Channels())
	wfx.NBlockAlign = uint16(audio.BlockAlign())
	wfx.NAvgBytesPerSec = uint32(audio.AvgBytesPerSec())
	wfx.CbSize = 0

	fmt.Println("--------")
	fmt.Printf("Format: PCM %d bit signed integer\n", wfx.WBitsPerSample)
	fmt.Printf("Rate: %d Hz\n", wfx.NSamplesPerSec)
	fmt.Printf("Channels: %d\n", wfx.NChannels)
	fmt.Println("--------")

	var defaultPeriodInFrames, fundamentalPeriodInFrames, minPeriodInFrames, maxPeriodInFrames uint32
	if err = ac3.GetSharedModeEnginePeriod(wfx, &defaultPeriodInFrames, &fundamentalPeriodInFrames, &minPeriodInFrames, &maxPeriodInFrames); err != nil {
		return
	}

	fmt.Println("Default period in frames: ", defaultPeriodInFrames)
	fmt.Println("Fundamental period in frames: ", fundamentalPeriodInFrames)
	fmt.Println("Min period in frames: ", minPeriodInFrames)
	fmt.Println("Max period in frames: ", maxPeriodInFrames)

	var latency time.Duration = time.Duration(float64(minPeriodInFrames)/float64(wfx.NSamplesPerSec)*1000) * time.Millisecond
	if err = ac3.InitializeSharedAudioStream(wca.AUDCLNT_SHAREMODE_SHARED, minPeriodInFrames, wfx, nil); err != nil {
		return
	}

	var bufferFrameSize uint32
	if err = ac3.GetBufferSize(&bufferFrameSize); err != nil {
		return
	}
	fmt.Println("Allocated buffer size: ", bufferFrameSize)
	fmt.Println("Latency: ", latency)
	fmt.Println("--------")

	var arc *wca.IAudioRenderClient
	if err = ac3.GetService(wca.IID_IAudioRenderClient, &arc); err != nil {
		return
	}
	defer arc.Release()

	if err = ac3.Start(); err != nil {
		return
	}

	fmt.Println("Start rendering audio with shared-timer-driven mode")
	fmt.Println("Press Ctrl-C to stop rendering")

	time.Sleep(latency)

	var input = audio.Bytes()
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
			if offset >= audio.Length() {
				isPlaying = false
				break
			}
			if err = ac3.GetCurrentPadding(&padding); err != nil {
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
			remaining := audio.Length() - offset
			if remaining < lim {
				lim = remaining
			}
			for n := 0; n < lim; n++ {
				b = (*byte)(unsafe.Pointer(uintptr(start) + uintptr(n)))
				*b = input[offset+n]
			}
			offset += lim
			if err = arc.ReleaseBuffer(availableFrameSize, 0); err != nil {
				return
			}
			// The buffer size is very small, we don't have to sleep for waiting render process.
			//time.Sleep(latency)
		}
	}
	time.Sleep(latency)
	return ac3.Stop()
}
