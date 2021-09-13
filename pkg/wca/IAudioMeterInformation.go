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
	GetPeakValue uintptr
}

func (v *IAudioMeterInformation) VTable() *IAudioMeterInformationVtbl {
	return (*IAudioMeterInformationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioMeterInformation) GetPeakValue(peak *float32) (err error) {
	err = amiGetPeakValue(v, peak)
	return
}
