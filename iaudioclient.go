package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioClient struct {
	ole.IUnknown
}

type IAudioClientVtbl struct {
	ole.IUnknownVtbl
	Initialize        uintptr
	GetBufferSize     uintptr
	GetStreamLatency  uintptr
	GetCurrentPadding uintptr
	IsFormatSupported uintptr
	GetMixFormat      uintptr
	GetDevicePeriod   uintptr
	Start             uintptr
	Stop              uintptr
	Reset             uintptr
	SetEventHandle    uintptr
	GetService        uintptr
}

func (v *IAudioClient) VTable() *IAudioClientVtbl {
	return (*IAudioClientVtbl)(unsafe.Pointer(v.RawVTable))
}
