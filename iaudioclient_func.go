// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func acInitialize(ac *IAudioClient, shareMode, streamFlags uint32, nsBufferDuration, nsPeriodicity REFERENCE_TIME, format *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetBufferSize(ac *IAudioClient, bufferFrameSize *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetStreamLatency(ac *IAudioClient, nsLatency *REFERENCE_TIME) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetCurrentPadding(ac *IAudioClient, numPadding *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acIsFormatSupported(ac *IAudioClient, shareMode uint32, wfx *WAVEFORMATEX, wfxClosestMatch **WAVEFORMATEX) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetMixFormat(ac *IAudioClient, wfx **WAVEFORMATEX) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetDevicePeriod(ac *IAudioClient, nsDefaultDevicePeriod, nsMinimumDevicePeriod *REFERENCE_TIME) (err error) {
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

func acSetEventHandle(ac *IAudioClient, handle uintptr) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acGetService(ac *IAudioClient, refIID *ole.GUID, obj interface{}) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
