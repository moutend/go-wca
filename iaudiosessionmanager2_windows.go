// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func asm2GetSessionEnumerator(asm2 *IAudioSessionManager2, sessionEnum **IAudioSessionEnumerator) (err error) {
	hr, _, _ := syscall.Syscall(
		asm2.VTable().GetSessionEnumerator,
		2,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionEnum)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2RegisterSessionNotification(asm2 *IAudioSessionManager2, sessionNotification *IAudioSessionNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		asm2.VTable().RegisterSessionNotification,
		2,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionNotification)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2UnregisterSessionNotification(asm2 *IAudioSessionManager2, sessionNotification *IAudioSessionNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		asm2.VTable().UnregisterSessionNotification,
		2,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionNotification)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2RegisterDuckNotification(asm2 *IAudioSessionManager2, sessionID *string, duckNotification *IAudioVolumeDuckNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		asm2.VTable().RegisterDuckNotification,
		3,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(sessionID)),
		uintptr(unsafe.Pointer(duckNotification)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asm2UnregisterDuckNotification(asm2 *IAudioSessionManager2, duckNotification *IAudioVolumeDuckNotification) (err error) {
	hr, _, _ := syscall.Syscall(
		asm2.VTable().UnregisterDuckNotification,
		2,
		uintptr(unsafe.Pointer(asm2)),
		uintptr(unsafe.Pointer(duckNotification)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
