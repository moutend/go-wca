//go:build !windows
// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func amiGetPeakValue(ami *IAudioMeterInformation, peak *float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func amiGetChannelsPeakValues(ami *IAudioMeterInformation, count uint32, peak []float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func amiGetMeteringChannelCount(ami *IAudioMeterInformation, count *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func amiQueryHardwareSupport(ami *IAudioMeterInformation, response *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
