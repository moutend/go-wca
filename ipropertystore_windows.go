// +build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func psGetCount(ps *IPropertyStore, count *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ps.VTable().GetCount,
		2,
		uintptr(unsafe.Pointer(ps)),
		uintptr(unsafe.Pointer(count)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getValue(ps *IPropertyStore, key *PropertyKey, pv *PROPVARIANT) (err error) {
	hr, _, _ := syscall.Syscall(
		ps.VTable().GetValue,
		3,
		uintptr(unsafe.Pointer(ps)),
		uintptr(unsafe.Pointer(key)),
		uintptr(unsafe.Pointer(&pv)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
