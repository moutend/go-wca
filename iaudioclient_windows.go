// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func acGetMixFormat(ac *IAudioClient, wfe **WAVEFORMATEX) (err error) {
	hr, _, _ := syscall.Syscall(
		ac.VTable().GetMixFormat,
		2,
		uintptr(unsafe.Pointer(ac)),
		uintptr(unsafe.Pointer(wfe)),
		0)
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
