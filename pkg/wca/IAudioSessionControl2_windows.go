//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func asc2GetSessionIdentifier(asc2 *IAudioSessionControl2, retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.SyscallN(
		asc2.VTable().GetSessionIdentifier,
		uintptr(unsafe.Pointer(asc2)),
		uintptr(unsafe.Pointer(&retValPtr)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	var us []uint16
	var i uint32
	var start = unsafe.Pointer(uintptr(retValPtr))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}
	*retVal = syscall.UTF16ToString(us)
	ole.CoTaskMemFree(uintptr(retValPtr))
	return
}

func asc2GetSessionInstanceIdentifier(asc2 *IAudioSessionControl2, retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.SyscallN(
		asc2.VTable().GetSessionInstanceIdentifier,
		uintptr(unsafe.Pointer(asc2)),
		uintptr(unsafe.Pointer(&retValPtr)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	var us []uint16
	var i uint32
	var start = unsafe.Pointer(uintptr(retValPtr))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}
	*retVal = syscall.UTF16ToString(us)
	ole.CoTaskMemFree(uintptr(retValPtr))
	return
}

func asc2GetProcessId(asc2 *IAudioSessionControl2, retVal *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc2.VTable().GetProcessId,
		uintptr(unsafe.Pointer(asc2)),
		uintptr(unsafe.Pointer(retVal)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asc2IsSystemSoundsSession(asc2 *IAudioSessionControl2) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc2.VTable().IsSystemSoundsSession,
		uintptr(unsafe.Pointer(asc2)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asc2SetDuckingPreference(asc2 *IAudioSessionControl2, optOut bool) (err error) {
	var optOutValue uint32

	if optOut {
		optOutValue = 1
	}
	hr, _, _ := syscall.SyscallN(
		asc2.VTable().SetDuckingPreference,
		uintptr(unsafe.Pointer(asc2)),
		uintptr(optOutValue))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
