// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmdcGetCount(dc *IMMDeviceCollection, count *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		dc.VTable().GetCount,
		2,
		uintptr(unsafe.Pointer(dc)),
		uintptr(unsafe.Pointer(count)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdcItem(dc *IMMDeviceCollection, id uint32, mmd **IMMDevice) (err error) {
	hr, _, _ := syscall.Syscall(
		dc.VTable().Item,
		3,
		uintptr(unsafe.Pointer(dc)),
		uintptr(id),
		uintptr(unsafe.Pointer(mmd)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
