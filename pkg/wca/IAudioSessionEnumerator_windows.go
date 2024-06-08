//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func aseGetCount(ase *IAudioSessionEnumerator, sessionCount *int) (err error) {
	hr, _, _ := syscall.SyscallN(
		ase.VTable().GetCount,
		uintptr(unsafe.Pointer(ase)),
		uintptr(unsafe.Pointer(sessionCount)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aseGetSession(ase *IAudioSessionEnumerator, sessionCount int, session **IAudioSessionControl) (err error) {
	hr, _, _ := syscall.SyscallN(
		ase.VTable().GetSession,
		uintptr(unsafe.Pointer(ase)),
		uintptr(sessionCount),
		uintptr(unsafe.Pointer(session)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
