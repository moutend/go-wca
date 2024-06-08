package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IAudioClient3 struct corresponds to IAudioClient3 interface.
// See MSDN's documentation:
// https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx
type IAudioClient3 struct {
	IAudioClient2
}

type IAudioClient3Vtbl struct {
	IAudioClient2Vtbl
	GetSharedModeEnginePeriod        uintptr
	GetCurrentSharedModeEnginePeriod uintptr
	InitializeSharedAudioStream      uintptr
}

func (v *IAudioClient3) VTable() *IAudioClient3Vtbl {
	return (*IAudioClient3Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioClient3) GetSharedModeEnginePeriod(wfx *WAVEFORMATEX, defaultPeriodInFrames, fundamentalPeriodInFrames, minPeriodInFrames, maxPeriodInFrames *uint32) (err error) {
	err = ac3GetSharedModeEnginePeriod(v, wfx, defaultPeriodInFrames, fundamentalPeriodInFrames, minPeriodInFrames, maxPeriodInFrames)
	return
}

func (v *IAudioClient3) GetCurrentSharedModeEnginePeriod(wfx **WAVEFORMATEX, currentPeriodInFrames *uint32) (err error) {
	err = ac3GetCurrentSharedModeEnginePeriod(v, wfx, currentPeriodInFrames)
	return
}

func (v *IAudioClient3) InitializeSharedAudioStream(streamFlags, periodInFrames uint32, wfx *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	err = ac3InitializeSharedAudioStream(v, streamFlags, periodInFrames, wfx, audioSessionGUID)
	return
}
