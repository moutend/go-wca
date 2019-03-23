package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioSessionControl struct {
	ole.IUnknown
}

type IAudioSessionControlVtbl struct {
	ole.IUnknownVtbl
	GetState                           uintptr
	GetDisplayName                     uintptr
	SetDisplayName                     uintptr
	GetIconPath                        uintptr
	SetIconPath                        uintptr
	GetGroupingParam                   uintptr
	SetGroupingParam                   uintptr
	RegisterAudioSessionNotification   uintptr
	UnregisterAudioSessionNotification uintptr
}

func (v *IAudioSessionControl) VTable() *IAudioSessionControlVtbl {
	return (*IAudioSessionControlVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionControl) GetState(retVal *uint32) (err error) {
	err = ascGetState(v, retVal)
	return
}

func (v *IAudioSessionControl) GetDisplayName(retVal *string) (err error) {
	err = ascGetDisplayName(v, retVal)
	return
}

func (v *IAudioSessionControl) SetDisplayName(value *string, eventContext *ole.GUID) (err error) {
	err = ascSetDisplayName(v, value, eventContext)
	return
}

func (v *IAudioSessionControl) GetIconPath(retVal *string) (err error) {
	err = ascGetIconPath(v, retVal)
	return
}

func (v *IAudioSessionControl) SetIconPath(value *string, eventContext *ole.GUID) (err error) {
	err = ascSetIconPath(v, value, eventContext)
	return
}

func (v *IAudioSessionControl) GetGroupingParam(retVal *ole.GUID) (err error) {
	err = ascGetGroupingParam(v, retVal)
	return
}

func (v *IAudioSessionControl) SetGroupingParam(override *ole.GUID, eventContext *ole.GUID) (err error) {
	err = ascSetGroupingParam(v, override, eventContext)
	return
}

func (v *IAudioSessionControl) RegisterAudioSessionNotification(newNotifications *IAudioSessionEvents) (err error) {
	err = ascRegisterAudioSessionNotification(v, newNotifications)
	return
}

func (v *IAudioSessionControl) UnregisterAudioSessionNotification(newNotifications *IAudioSessionEvents) (err error) {
	err = ascUnregisterAudioSessionNotification(v, newNotifications)
	return
}
