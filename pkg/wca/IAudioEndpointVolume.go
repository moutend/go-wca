package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioEndpointVolume struct {
	ole.IUnknown
}

type IAudioEndpointVolumeVtbl struct {
	ole.IUnknownVtbl
	RegisterControlChangeNotify   uintptr
	UnregisterControlChangeNotify uintptr
	GetChannelCount               uintptr
	SetMasterVolumeLevel          uintptr
	SetMasterVolumeLevelScalar    uintptr
	GetMasterVolumeLevel          uintptr
	GetMasterVolumeLevelScalar    uintptr
	SetChannelVolumeLevel         uintptr
	SetChannelVolumeLevelScalar   uintptr
	GetChannelVolumeLevel         uintptr
	GetChannelVolumeLevelScalar   uintptr
	SetMute                       uintptr
	GetMute                       uintptr
	GetVolumeStepInfo             uintptr
	VolumeStepUp                  uintptr
	VolumeStepDown                uintptr
	QueryHardwareSupport          uintptr
	GetVolumeRange                uintptr
}

func (v *IAudioEndpointVolume) VTable() *IAudioEndpointVolumeVtbl {
	return (*IAudioEndpointVolumeVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioEndpointVolume) RegisterControlChangeNotify() (err error) {
	err = aevRegisterControlChangeNotify()
	return
}

func (v *IAudioEndpointVolume) UnregisterControlChangeNotify() (err error) {
	err = aevUnregisterControlChangeNotify()
	return
}

func (v *IAudioEndpointVolume) GetChannelCount(channelCount *uint32) (err error) {
	err = aevGetChannelCount(v, channelCount)
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevel(levelDB float32, eventContextGUID *ole.GUID) (err error) {
	err = aevSetMasterVolumeLevel(v, levelDB, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevelScalar(level float32, eventContextGUID *ole.GUID) (err error) {
	err = aevSetMasterVolumeLevelScalar(v, level, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) GetMasterVolumeLevel(level *float32) (err error) {
	err = aevGetMasterVolumeLevel(v, level)
	return
}

func (v *IAudioEndpointVolume) GetMasterVolumeLevelScalar(level *float32) (err error) {
	err = aevGetMasterVolumeLevelScalar(v, level)
	return
}

func (v *IAudioEndpointVolume) SetChannelVolumeLevel(channel uint32, levelDB float32, eventContextGUID *ole.GUID) (err error) {
	err = aevSetChannelVolumeLevel(v, channel, levelDB, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) SetChannelVolumeLevelScalar(channel uint32, level float32, eventContextGUID *ole.GUID) (err error) {
	err = aevSetChannelVolumeLevelScalar(v, channel, level, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) GetChannelVolumeLevel(channel uint32, levelDB *float32) (err error) {
	err = aevGetChannelVolumeLevel(v, channel, levelDB)
	return
}

func (v *IAudioEndpointVolume) GetChannelVolumeLevelScalar(channel uint32, level *float32) (err error) {
	err = aevGetChannelVolumeLevelScalar(v, channel, level)
	return
}

func (v *IAudioEndpointVolume) SetMute(mute bool, eventContextGUID *ole.GUID) (err error) {
	err = aevSetMute(v, mute, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) GetMute(mute *bool) (err error) {
	err = aevGetMute(v, mute)
	return
}

func (v *IAudioEndpointVolume) GetVolumeStepInfo(step, stepCount *uint32) (err error) {
	err = aevGetVolumeStepInfo(v, step, stepCount)
	return
}

func (v *IAudioEndpointVolume) VolumeStepUp(eventContextGUID *ole.GUID) (err error) {
	err = aevVolumeStepUp(v, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) VolumeStepDown(eventContextGUID *ole.GUID) (err error) {
	err = aevVolumeStepDown(v, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) QueryHardwareSupport(hardwareSupportMask *uint32) (err error) {
	err = aevQueryHardwareSupport(v, hardwareSupportMask)
	return
}

func (v *IAudioEndpointVolume) GetVolumeRange(minDB, maxDB, incrementDB *float32) (err error) {
	err = aevGetVolumeRange(v, minDB, maxDB, incrementDB)
	return
}
