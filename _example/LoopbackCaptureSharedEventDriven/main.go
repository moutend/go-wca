package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wav"
	"github.com/moutend/go-wca/pkg/wca"
)

var version = "latest"
var revision = "latest"

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
	var versionFlag bool
	var audio *wav.File
	var file []byte

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Var(&durationFlag, "duration", "Specify recording duration in second")
	f.Var(&durationFlag, "d", "Alias of --duration")
	f.Var(&filenameFlag, "output", "file name")
	f.Var(&filenameFlag, "o", "Alias of --output")
	f.BoolVar(&versionFlag, "version", false, "Show version")
	f.Parse(args[1:])

	if versionFlag {
		fmt.Printf("%s-%s\n", version, revision)
		return
	}
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
	if file, err = wav.Marshal(audio); err != nil {
		return
	}
	if err = ioutil.WriteFile(filenameFlag.Value, file, 0644); err != nil {
		return
	}
	fmt.Println("Successfully done")
	return
}

func loopbackCaptureSharedEventDriven(ctx context.Context, duration time.Duration) (audio *wav.File, err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}
	defer ole.CoUninitialize()

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

	var cac *wca.IAudioClient
	if err = mmdCapture.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &cac); err != nil {
		return
	}
	defer cac.Release()

	var rac *wca.IAudioClient
	if err = mmdRender.Activate(wca.IID_IAudioClient, wca.CLSCTX_ALL, nil, &rac); err != nil {
		return
	}
	defer rac.Release()

	var wfx *wca.WAVEFORMATEX
	if err = rac.GetMixFormat(&wfx); err != nil {
		return
	}
	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(wfx)))

	wfx.WFormatTag = 1
	wfx.NBlockAlign = (wfx.WBitsPerSample / 8) * wfx.NChannels
	wfx.NAvgBytesPerSec = wfx.NSamplesPerSec * uint32(wfx.NBlockAlign)
	wfx.CbSize = 0

	if audio, err = wav.New(int(wfx.NSamplesPerSec), int(wfx.WBitsPerSample), int(wfx.NChannels)); err != nil {
		return
	}

	fmt.Println("--------")
	fmt.Printf("Format: PCM %d bit signed integer\n", wfx.WBitsPerSample)
	fmt.Printf("Rate: %d Hz\n", wfx.NSamplesPerSec)
	fmt.Printf("Channels: %d\n", wfx.NChannels)
	fmt.Println("--------")

	var defaultPeriod wca.REFERENCE_TIME
	var minimumPeriod wca.REFERENCE_TIME
	var latency time.Duration
	if err = rac.GetDevicePeriod(&defaultPeriod, &minimumPeriod); err != nil {
		return
	}
	latency = time.Duration(int(defaultPeriod) * 100)

	fmt.Println("Default period: ", defaultPeriod)
	fmt.Println("Minimum period: ", minimumPeriod)
	fmt.Println("Latency: ", latency)

	if err = cac.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, wca.AUDCLNT_STREAMFLAGS_EVENTCALLBACK|wca.AUDCLNT_STREAMFLAGS_LOOPBACK, defaultPeriod, 0, wfx, nil); err != nil {
		return
	}
	if err = rac.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, wca.AUDCLNT_STREAMFLAGS_EVENTCALLBACK, defaultPeriod, 0, wfx, nil); err != nil {
		return
	}

	fakeAudioReadyEvent := wca.CreateEventExA(0, 0, 0, wca.EVENT_MODIFY_STATE|wca.SYNCHRONIZE)
	defer wca.CloseHandle(fakeAudioReadyEvent)

	if err = cac.SetEventHandle(fakeAudioReadyEvent); err != nil {
		return
	}

	audioReadyEvent := wca.CreateEventExA(0, 0, 0, wca.EVENT_MODIFY_STATE|wca.SYNCHRONIZE)
	defer wca.CloseHandle(audioReadyEvent)

	if err = rac.SetEventHandle(audioReadyEvent); err != nil {
		return
	}

	var bufferFrameSizeRender uint32
	if err = rac.GetBufferSize(&bufferFrameSizeRender); err != nil {
		return
	}

	var bufferFrameSize uint32
	if err = cac.GetBufferSize(&bufferFrameSize); err != nil {
		return
	}

	fmt.Printf("Allocated buffer size: %d\n", bufferFrameSize)

	var arc *wca.IAudioRenderClient
	if err = rac.GetService(wca.IID_IAudioRenderClient, &arc); err != nil {
		return
	}
	defer arc.Release()

	var acc *wca.IAudioCaptureClient
	if err = cac.GetService(wca.IID_IAudioCaptureClient, &acc); err != nil {
		return
	}
	defer acc.Release()

	if err = rac.Start(); err != nil {
		return
	}
	if err = cac.Start(); err != nil {
		return
	}

	fmt.Println("Start loopback capturing with shared event driven mode")

	if duration <= 0 {
		fmt.Println("Press Ctrl-C to save and quit")
	}

	var output = []byte{}
	var buf []byte
	var offset int
	var lim int
	var start unsafe.Pointer
	var isCapturing bool = true
	var currentDuration time.Duration
	var data *byte
	var b *byte
	var availableFrameSize uint32
	var flags uint32
	var devicePosition uint64
	var qcpPosition uint64

	errorChan := make(chan error, 1)

	time.Sleep(latency)

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
			currentDuration = time.Duration(float64(offset) / float64(wfx.WBitsPerSample/8) / float64(wfx.NChannels) / float64(wfx.NSamplesPerSec) * float64(time.Second))
			if duration != 0 && currentDuration > duration {
				isCapturing = false
				break
			}
			if err != nil {
				isCapturing = false
				break
			}
			if err = acc.GetBuffer(&data, &availableFrameSize, &flags, &devicePosition, &qcpPosition); err != nil {
				continue
			}
			if availableFrameSize == 0 {
				continue
			}

			start = unsafe.Pointer(data)
			lim = int(availableFrameSize) * int(wfx.NBlockAlign)
			buf = make([]byte, lim)

			for n := 0; n < lim; n++ {
				b = (*byte)(unsafe.Pointer(uintptr(start) + uintptr(n)))
				buf[n] = *b
			}

			offset += lim
			output = append(output, buf...)

			if err = acc.ReleaseBuffer(availableFrameSize); err != nil {
				return
			}
		}
	}

	io.Copy(audio, bytes.NewBuffer(output))

	fmt.Println("Stop capturing")
	if err = cac.Stop(); err != nil {
		return
	}
	if err = rac.Stop(); err != nil {
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
