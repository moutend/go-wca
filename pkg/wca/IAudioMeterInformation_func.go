// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func amiGetPeakValue(ami *IAudioMeterInformation, peak *float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func amiGetChannelsPeakValues(ami *IAudioMeterInformation, peak *float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func amiGetMeteringChannelCount(ami *IAudioMeterInformation, peak *float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func amiQueryHardwareSupport(ami *IAudioMeterInformation, peak *float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
