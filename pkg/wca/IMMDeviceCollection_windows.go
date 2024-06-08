//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmdcGetCount(dc *IMMDeviceCollection, count *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		dc.VTable().GetCount,
		uintptr(unsafe.Pointer(dc)),
		uintptr(unsafe.Pointer(count)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdcItem(dc *IMMDeviceCollection, id uint32, mmd **IMMDevice) (err error) {
	hr, _, _ := syscall.SyscallN(
		dc.VTable().Item,
		uintptr(unsafe.Pointer(dc)),
		uintptr(id),
		uintptr(unsafe.Pointer(mmd)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
