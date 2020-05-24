// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func ac3GetSharedModeEnginePeriod(ac3 *IAudioClient3, wfx *WAVEFORMATEX, defaultPeriodInFrames, fundamentalPeriodInFrames, minPeriodInFrames, maxPeriodInFrames *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ac3GetCurrentSharedModeEnginePeriod(ac3 *IAudioClient3, wfx **WAVEFORMATEX, currentPeriodInFrames *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ac3InitializeSharedAudioStream(ac3 *IAudioClient3, streamFlags, periodInFrames uint32, wfx *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
