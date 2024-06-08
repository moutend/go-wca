//go:build windows

package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func acInitialize(ac *IAudioClient, shareMode, streamFlags uint32, nsBufferDuration, nsPeriodicity REFERENCE_TIME, format *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().Initialize,
		uintptr(unsafe.Pointer(ac)),
		uintptr(shareMode),
		uintptr(streamFlags),
		uintptr(nsBufferDuration),
		uintptr(nsPeriodicity),
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(audioSessionGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acInitializeEx(ac *IAudioClient, shareMode, streamFlags uint32, nsBufferDuration, nsPeriodicity REFERENCE_TIME, format *WAVEFORMATEXTENSIBLE, audioSessionGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().Initialize,
		uintptr(unsafe.Pointer(ac)),
		uintptr(shareMode),
		uintptr(streamFlags),
		uintptr(nsBufferDuration),
		uintptr(nsPeriodicity),
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(audioSessionGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetBufferSize(ac *IAudioClient, bufferFrameSize *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetBufferSize,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(bufferFrameSize)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetStreamLatency(ac *IAudioClient, nsLatency *REFERENCE_TIME) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetStreamLatency,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(nsLatency)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetCurrentPadding(ac *IAudioClient, numPadding *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetCurrentPadding,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(numPadding)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acIsFormatSupported(ac *IAudioClient, shareMode uint32, wfx *WAVEFORMATEX, wfxClosestMatch **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().IsFormatSupported,
		uintptr(unsafe.Pointer(ac)),
		uintptr(shareMode),
		uintptr(unsafe.Pointer(wfx)),
		uintptr(unsafe.Pointer(wfxClosestMatch)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetMixFormat(ac *IAudioClient, wfx **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetMixFormat,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(wfx)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetMixFormatEx(ac *IAudioClient, wfe **WAVEFORMATEXTENSIBLE) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetMixFormat,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(wfe)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetDevicePeriod(ac *IAudioClient, nsDefaultDevicePeriod, nsMinimumDevicePeriod *REFERENCE_TIME) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetDevicePeriod,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(nsDefaultDevicePeriod)),
		uintptr(unsafe.Pointer(nsMinimumDevicePeriod)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acStart(ac *IAudioClient) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().Start,
		uintptr(unsafe.Pointer(ac)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acStop(ac *IAudioClient) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().Stop,
		uintptr(unsafe.Pointer(ac)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acReset(ac *IAudioClient) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().Reset,
		uintptr(unsafe.Pointer(ac)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acSetEventHandle(ac *IAudioClient, handle uintptr) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac.VTable().SetEventHandle,
		uintptr(unsafe.Pointer(ac)),
		handle)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func acGetService(ac *IAudioClient, refIID *ole.GUID, obj interface{}) (err error) {
	objValue := reflect.ValueOf(obj).Elem()
	hr, _, _ := syscall.SyscallN(
		ac.VTable().GetService,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(refIID)),
		objValue.Addr().Pointer())
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
