package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioSessionManager struct {
	ole.IUnknown
}

type IAudioSessionManagerVtbl struct {
	ole.IUnknownVtbl
	GetAudioSessionControl uintptr
	GetSimpleAudioVolume   uintptr
}

func (v *IAudioSessionManager) VTable() *IAudioSessionManagerVtbl {
	return (*IAudioSessionManagerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionManager) GetAudioSessionControl(audioSessionGUID *ole.GUID, streamFlags uint32, sessionControl **IAudioSessionControl) (err error) {
	err = asmGetAudioSessionControl(v, audioSessionGUID, streamFlags, sessionControl)
	return
}

func (v *IAudioSessionManager) GetSimpleAudioVolume(audioSessionGUID *ole.GUID, streamFlags uint32, audioVolume **ISimpleAudioVolume) (err error) {
	err = asmGetSimpleAudioVolume(v, audioSessionGUID, streamFlags, audioVolume)
	return
}
