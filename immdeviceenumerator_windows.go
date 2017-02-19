// +build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func enumAudioEndpoints(de *IMMDeviceEnumerator, eDataFlow int32, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	hr, _, _ := syscall.Syscall6(
		de.VTable().EnumAudioEndpoints,
		4,
		uintptr(unsafe.Pointer(de)),
		0, //uintptr(unsafe.Pointer(&eDataFlow)),
		1, //uintptr(unsafe.Pointer(&stateMask)),
		uintptr(unsafe.Pointer(dc)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
