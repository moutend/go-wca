package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type ISimpleAudioVolume struct {
	ole.IUnknown
}

type ISimpleAudioVolumeVtbl struct {
	ole.IUnknownVtbl
	SetMasterVolume uintptr
	GetMasterVolume uintptr
	SetMute         uintptr
	GetMute         uintptr
}

func (v *ISimpleAudioVolume) VTable() *ISimpleAudioVolumeVtbl {
	return (*ISimpleAudioVolumeVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ISimpleAudioVolume) SetMasterVolume(level float32, eventContext *ole.GUID) (err error) {
	err = savSetMasterVolume(v, level, eventContext)
	return
}

func (v *ISimpleAudioVolume) GetMasterVolume(level *float32) (err error) {
	err = savGetMasterVolume(v, level)
	return
}

func (v *ISimpleAudioVolume) SetMute(mute bool, eventContext *ole.GUID) (err error) {
	err = savSetMute(v, mute, eventContext)
	return
}

func (v *ISimpleAudioVolume) GetMute(mute *bool) (err error) {
	err = savGetMute(v, mute)
	return
}
