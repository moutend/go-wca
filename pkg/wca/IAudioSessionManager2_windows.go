//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func asm2GetSessionEnumerator(asm2 *IAudioSessionManager2, sessionEnum **IAudioSessionEnumerator) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm2.VTable().GetSessionEnumerator,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionEnum)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2RegisterSessionNotification(asm2 *IAudioSessionManager2, sessionNotification *IAudioSessionNotification) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm2.VTable().RegisterSessionNotification,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionNotification)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2UnregisterSessionNotification(asm2 *IAudioSessionManager2, sessionNotification *IAudioSessionNotification) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm2.VTable().UnregisterSessionNotification,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionNotification)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2RegisterDuckNotification(asm2 *IAudioSessionManager2, sessionID *string, duckNotification *IAudioVolumeDuckNotification) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm2.VTable().RegisterDuckNotification,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionID)),
		uintptr(unsafe.Pointer(duckNotification)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2UnregisterDuckNotification(asm2 *IAudioSessionManager2, duckNotification *IAudioVolumeDuckNotification) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm2.VTable().UnregisterDuckNotification,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(duckNotification)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
