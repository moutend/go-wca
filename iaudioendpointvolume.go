package main

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

func (v *IAudioEndpointVolume) RegisterControlChangeNotify() {
	return
}
func (v *IAudioEndpointVolume) UnregisterControlChangeNotify() {
	return
}

func (v *IAudioEndpointVolume) GetChannelCount(channelCount *uint32) (err error) {
	err = getChannelCount(v, channelCount)
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevel() {
	return
}

func (v *IAudioEndpointVolume) SetMasterVolumeLevelScalar(level float32, eventContextGUID *ole.GUID) (err error) {
	err = setMasterVolumeLevelScalar(v, level, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) GetMasterVolumeLevel() {
	return
}
func (v *IAudioEndpointVolume) GetMasterVolumeLevelScalar(level *float32) (err error) {
	err = getMasterVolumeLevelScalar(v, level)
	return
}
func (v *IAudioEndpointVolume) SetChannelVolumeLevel() {
	return
}
func (v *IAudioEndpointVolume) SetChannelVolumeLevelScalar() {
	return
}
func (v *IAudioEndpointVolume) GetChannelVolumeLevel() {
	return
}
func (v *IAudioEndpointVolume) GetChannelVolumeLevelScalar() {
	return
}
func (v *IAudioEndpointVolume) SetMute() {
	return
}
func (v *IAudioEndpointVolume) GetMute() {
	return
}
func (v *IAudioEndpointVolume) GetVolumeStepInfo(step, stepCount *uint32) (err error) {
	err = getVolumeStepInfo(v, step, stepCount)
	return
}

func (v *IAudioEndpointVolume) VolumeStepUp(eventContextGUID *ole.GUID) (err error) {
	err = volumeStepUp(v, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) VolumeStepDown(eventContextGUID *ole.GUID) (err error) {
	err = volumeStepDown(v, eventContextGUID)
	return
}

func (v *IAudioEndpointVolume) QueryHardwareSupport(hardwareSupportMask *uint32) (err error) {
	err = queryHardwareSupport(v, hardwareSupportMask)
	return
}

func (v *IAudioEndpointVolume) GetVolumeRange(minDB, maxDB, incrementDB *float32) (err error) {
	err = getVolumeRange(v, minDB, maxDB, incrementDB)
	return
}
