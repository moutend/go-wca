// +build windows
package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/gwca"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully done")
	}
	return
}

func run() (err error) {
	output, err := capture()
	if err = ioutil.WriteFile("output.raw", output, 0644); err != nil {
		return
	}
	fmt.Println("Saved captured audio as output.raw")
	return
}

func capture() (output []byte, err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}

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
	fmt.Printf("Capturing what you hear from\n%s\n", pv.String())

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
		wfx.WBitsPerSample = 16
		wfx.NBlockAlign = (wfx.WBitsPerSample / 8) * wfx.NChannels // 16 bit stereo is 32bit (4 byte) per sample
		wfx.NAvgBytesPerSec = wfx.NSamplesPerSec * uint32(wfx.NBlockAlign)
		wfx.CbSize = 0
	}

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
	capturingPeriod = time.Duration(defaultPeriod * 100)
	fmt.Printf("Default capturing period: %d ms\n", capturingPeriod/time.Millisecond)

	if err = ac.Initialize(wca.AUDCLNT_SHAREMODE_SHARED, wca.AUDCLNT_STREAMFLAGS_LOOPBACK, 500*10000, 0, wfx, nil); err != nil {
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

	var data *byte
	for m := 0; m < 2000; m++ {
		var availableFrameSize uint32
		var flags uint32
		var devicePosition uint64
		var qcpPosition uint64
		if err = acc.GetBuffer(&data, &availableFrameSize, &flags, &devicePosition, &qcpPosition); err != nil {
			return
		}

		start := unsafe.Pointer(data)
		for n := 0; n < int(availableFrameSize)*int(wfx.NBlockAlign); n++ {
			var b *byte
			b = (*byte)(unsafe.Pointer(uintptr(start) + uintptr(n)))
			output = append(output, *b)
		}
		time.Sleep(capturingPeriod)
		if err = acc.ReleaseBuffer(availableFrameSize); err != nil {
			return
		}
	}

	if err = ac.Stop(); err != nil {
		return
	}
	fmt.Println("Stopping capturing loopback audio")
	return
}
