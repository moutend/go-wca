package wca

import (
	"unsafe"
)

type IAudioSessionControl2 struct {
	IAudioSessionControl
}

type IAudioSessionControl2Vtbl struct {
	IAudioSessionControlVtbl
	GetSessionIdentifier         uintptr
	GetSessionInstanceIdentifier uintptr
	GetProcessId                 uintptr
	IsSystemSoundsSession        uintptr
	SetDuckingPreference         uintptr
}

func (v *IAudioSessionControl2) VTable() *IAudioSessionControl2Vtbl {
	return (*IAudioSessionControl2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionControl2) GetSessionIdentifier(retVal *string) (err error) {
	err = asc2GetSessionIdentifier(v, retVal)
	return
}

func (v *IAudioSessionControl2) GetSessionInstanceIdentifier(retVal *string) (err error) {
	err = asc2GetSessionInstanceIdentifier(v, retVal)
	return
}

func (v *IAudioSessionControl2) GetProcessId(retVal *uint32) (err error) {
	err = asc2GetProcessId(v, retVal)
	return
}

func (v *IAudioSessionControl2) IsSystemSoundsSession() (err error) {
	err = asc2IsSystemSoundsSession(v)
	return
}

func (v *IAudioSessionControl2) SetDuckingPreference(optOut bool) (err error) {
	err = asc2SetDuckingPreference(v, optOut)
	return
}
