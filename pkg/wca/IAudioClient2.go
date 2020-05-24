package wca

import (
	"unsafe"
)

// IAudioClient2 struct corresponds to IAudioClient2 interface.
// See MSDN's documentation:
// https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx
type IAudioClient2 struct {
	IAudioClient
}

type IAudioClient2Vtbl struct {
	IAudioClientVtbl
	IsOffloadCapable    uintptr
	SetClientProperties uintptr
	GetBufferSizeLimits uintptr
}

func (v *IAudioClient2) VTable() *IAudioClient2Vtbl {
	return (*IAudioClient2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioClient2) IsOffloadCapable(category uint32, isOffloadCapable *bool) (err error) {
	err = ac2IsOffloadCapable(v, category, isOffloadCapable)
	return
}

func (v *IAudioClient2) SetClientProperties(properties *AudioClientProperties) (err error) {
	err = ac2SetClientProperties(v, properties)
	return
}

func (v *IAudioClient2) GetBufferSizeLimits(wfx *WAVEFORMATEX, isEventDriven bool, minBufferDuration, maxBufferDuration *uint32) (err error) {
	err = ac2GetBufferSizeLimits(v, wfx, isEventDriven, minBufferDuration, maxBufferDuration)
	return
}
