//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func psGetCount(ps *IPropertyStore, count *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		ps.VTable().GetCount,
		uintptr(unsafe.Pointer(ps)),
		uintptr(unsafe.Pointer(count)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func psGetAt(ps *IPropertyStore, iProp uint32, pkey *PROPERTYKEY) (err error) {
	hr, _, _ := syscall.SyscallN(
		ps.VTable().GetAt,
		uintptr(unsafe.Pointer(ps)),
		uintptr(iProp),
		uintptr(unsafe.Pointer(pkey)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func psGetValue(ps *IPropertyStore, key *PROPERTYKEY, pv *PROPVARIANT) (err error) {
	hr, _, _ := syscall.SyscallN(
		ps.VTable().GetValue,
		uintptr(unsafe.Pointer(ps)),
		uintptr(unsafe.Pointer(key)),
		uintptr(unsafe.Pointer(pv)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func psSetValue() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psCommit() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
