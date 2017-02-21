// +build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func getId(mmd *IMMDevice, strId *uint16) (err error) {
	hr, _, _ := syscall.Syscall(
		mmd.VTable().GetId,
		2,
		uintptr(unsafe.Pointer(mmd)),
		uintptr(unsafe.Pointer(strId)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
