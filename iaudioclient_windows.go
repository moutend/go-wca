// +build windows

package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func acInitialize(ac *IAudioClient, shareMode, streamFlags uint32, nsBufferDuration, nsPeriodicity REFERENCE_TIME, format *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall9(
		ac.VTable().Initialize,
		7,
		uintptr(unsafe.Pointer(ac)),
		uintptr(shareMode),
		uintptr(streamFlags),
		uintptr(nsBufferDuration),
		uintptr(nsPeriodicity),
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(audioSessionGUID)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetBufferSize(ac *IAudioClient, bufferFrameSize *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetBufferSize,
		2,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(bufferFrameSize)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetStreamLatency(ac *IAudioClient, nsLatency *REFERENCE_TIME) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetStreamLatency,
		2,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(nsLatency)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetCurrentPadding(ac *IAudioClient, numPadding *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetCurrentPadding,
		2,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(numPadding)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acIsFormatSupported(ac *IAudioClient, shareMode uint32, wfx *WAVEFORMATEX, wfxClosestMatch **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.Syscall6(
		ac.VTable().IsFormatSupported,
		4,
		uintptr(unsafe.Pointer(ac)),
		uintptr(shareMode),
		uintptr(unsafe.Pointer(wfx)),
		uintptr(unsafe.Pointer(wfxClosestMatch)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetMixFormat(ac *IAudioClient, wfx **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetMixFormat,
		2,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(wfx)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetDevicePeriod(ac *IAudioClient, nsDefaultDevicePeriod, nsMinimumDevicePeriod *REFERENCE_TIME) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetDevicePeriod,
		3,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(nsDefaultDevicePeriod)),
		uintptr(unsafe.Pointer(nsMinimumDevicePeriod)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acStart(ac *IAudioClient) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().Start,
		1,
		uintptr(unsafe.Pointer(ac)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acStop(ac *IAudioClient) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().Stop,
		1,
		uintptr(unsafe.Pointer(ac)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acReset(ac *IAudioClient) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().Reset,
		1,
		uintptr(unsafe.Pointer(ac)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acSetEventHandle(ac *IAudioClient, handle uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().SetEventHandle,
		2,
		uintptr(unsafe.Pointer(ac)),
		uintptr(handle),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetService(ac *IAudioClient, refIID *ole.GUID, obj interface{}) (err error) {
	objValue := reflect.ValueOf(obj).Elem()
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetService,
		3,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(refIID)),
		objValue.Addr().Pointer())
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
