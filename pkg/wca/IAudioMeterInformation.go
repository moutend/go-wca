package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioMeterInformation struct {
	ole.IUnknown
}

type IAudioMeterInformationVtbl struct {
	ole.IUnknownVtbl
	GetPeakValue            uintptr
	GetChannelsPeakValues   uintptr
	GetMeteringChannelCount uintptr
	QueryHardwareSupport    uintptr
}

func (v *IAudioMeterInformation) VTable() *IAudioMeterInformationVtbl {
	return (*IAudioMeterInformationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioMeterInformation) GetPeakValue(peak *float32) (err error) {
	err = amiGetPeakValue(v, peak)
	return
}

func (v *IAudioMeterInformation) GetMeteringChannelCount(count *uint32) (err error) {
	err = amiGetMeteringChannelCount(v, count)
	return
}

func (v *IAudioMeterInformation) GetChannelsPeakValues(count uint32, peaks []float32) (err error) {
	err = amiGetChannelsPeakValues(v, count, peaks)
	return
}

func (v *IAudioMeterInformation) QueryHardwareSupport(response *uint32) (err error) {
	err = amiQueryHardwareSupport(v, response)
	return
}
