// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func acInitialize(ac *IAudioClient, shareMode, streamFlags, bufferDuration, periodicity uint32, format *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetBufferSize(ac *IAudioClient, bufferFrameSize *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetMixFormat(ac *IAudioClient, wfe **WAVEFORMATEX) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acStart(ac *IAudioClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acStop(ac *IAudioClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acReset(ac *IAudioClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
