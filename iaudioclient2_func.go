// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func ac2IsOffloadCapable(ac2 *IAudioClient2, category uint32, isOffloadCapable *bool) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ac2SetClientProperties(ac2 *IAudioClient2, properties *AudioClientProperties) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ac2GetBufferSizeLimits(ac2 *IAudioClient2, wfx *WAVEFORMATEX, isEventDriven bool, minBufferDuration, maxBufferDuration *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
