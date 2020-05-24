package wca

import (
	"unsafe"
)

type IAudioSessionManager2 struct {
	IAudioSessionManager
}

type IAudioSessionManager2Vtbl struct {
	IAudioSessionManagerVtbl
	GetSessionEnumerator          uintptr
	RegisterSessionNotification   uintptr
	UnregisterSessionNotification uintptr
	RegisterDuckNotification      uintptr
	UnregisterDuckNotification    uintptr
}

func (v *IAudioSessionManager2) VTable() *IAudioSessionManager2Vtbl {
	return (*IAudioSessionManager2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionManager2) GetSessionEnumerator(sessionEnum **IAudioSessionEnumerator) (err error) {
	err = asm2GetSessionEnumerator(v, sessionEnum)
	return
}

func (v *IAudioSessionManager2) RegisterSessionNotification(sessionNotification *IAudioSessionNotification) (err error) {
	err = asm2RegisterSessionNotification(v, sessionNotification)
	return
}

func (v *IAudioSessionManager2) UnregisterSessionNotification(sessionNotification *IAudioSessionNotification) (err error) {
	err = asm2UnregisterSessionNotification(v, sessionNotification)
	return
}

func (v *IAudioSessionManager2) RegisterDuckNotification(sessionID *string, duckNotification *IAudioVolumeDuckNotification) (err error) {
	err = asm2RegisterDuckNotification(v, sessionID, duckNotification)
	return
}

func (v *IAudioSessionManager2) UnregisterDuckNotification(duckNotification *IAudioVolumeDuckNotification) (err error) {
	err = asm2UnregisterDuckNotification(v, duckNotification)
	return
}
