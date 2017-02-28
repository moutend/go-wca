// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmdeEnumAudioEndpoints(de *IMMDeviceEnumerator, eDataFlow, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	hr, _, _ := syscall.Syscall6(
		de.VTable().EnumAudioEndpoints,
		4,
		uintptr(unsafe.Pointer(de)),
		uintptr(eDataFlow), //uintptr(unsafe.Pointer(&eDataFlow)),
		uintptr(stateMask), //uintptr(unsafe.Pointer(&stateMask)),
		uintptr(unsafe.Pointer(dc)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdeGetDefaultAudioEndpoint() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeGetDevice() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeRegisterEndpointNotificationCallback() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeUnregisterEndpointNotificationCallback() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
