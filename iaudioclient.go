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

func (v *IAudioClient) GetMixFormat(wfe **WAVEFORMATEX) (err error) {
	err = acGetMixFormat(v, wfe)
	return
}

func (v *IAudioClient) Start() (err error) {
	err = acStart(v)
	return
}

func (v *IAudioClient) Stop() (err error) {
	err = acStop(v)
	return
}

func (v *IAudioClient) Reset() (err error) {
	err = acReset(v)
	return
}
