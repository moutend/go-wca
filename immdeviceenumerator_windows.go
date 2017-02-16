// +build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func enumAudioEndpoints(de *IMMDeviceEnumerator, eDataFlow uint32, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	hr, _, _ := syscall.Syscall6(
		de.VTable().EnumAudioEndpoints,
		4,
		uintptr(unsafe.Pointer(de)),
		uintptr(unsafe.Pointer(&eDataFlow)),
		uintptr(unsafe.Pointer(&stateMask)),
		uintptr(unsafe.Pointer(dc)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
