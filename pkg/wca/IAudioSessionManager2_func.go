// +build !windows

package wca

import "github.com/go-ole/go-ole"

func asm2GetSessionEnumerator(asm2 *IAudioSessionManager2, sessionEnum **IAudioSessionEnumerator) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asm2RegisterSessionNotification(asm2 *IAudioSessionManager2, sessionNotification *IAudioSessionNotification) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asm2UnregisterSessionNotification(asm2 *IAudioSessionManager2, sessionNotification *IAudioSessionNotification) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asm2RegisterDuckNotification(asm2 *IAudioSessionManager2, sessionID *string, duckNotification *IAudioVolumeDuckNotification) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asm2UnregisterDuckNotification(asm2 *IAudioSessionManager2, duckNotification *IAudioVolumeDuckNotification) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
