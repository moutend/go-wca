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

func (v *IAudioClient) Initialize(shareMode, streamFlags uint32, nsBufferDuration, nsPeriodicity REFERENCE_TIME, format *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	err = acInitialize(v, shareMode, streamFlags, nsBufferDuration, nsPeriodicity, format, audioSessionGUID)
	return
}

func (v *IAudioClient) GetBufferSize(bufferFrameSize *uint32) (err error) {
	err = acGetBufferSize(v, bufferFrameSize)
	return
}

func (v *IAudioClient) GetStreamLatency(nsLatency *REFERENCE_TIME) (err error) {
	err = acGetStreamLatency(v, nsLatency)
	return
}

func (v *IAudioClient) GetCurrentPadding(numPadding *uint32) (err error) {
	err = acGetCurrentPadding(v, numPadding)
	return
}

func (v *IAudioClient) IsFormatSupported(shareMode uint32, wfx *WAVEFORMATEX, wfxClosestMatch **WAVEFORMATEX) (err error) {
	err = acIsFormatSupported(v, shareMode, wfx, wfxClosestMatch)
	return
}
func (v *IAudioClient) GetMixFormat(wfx **WAVEFORMATEX) (err error) {
	err = acGetMixFormat(v, wfx)
	return
}

func (v *IAudioClient) GetDevicePeriod(nsDefaultDevicePeriod, nsMinimumDevicePeriod *REFERENCE_TIME) (err error) {
	err = acGetDevicePeriod(v, nsDefaultDevicePeriod, nsMinimumDevicePeriod)
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

func (v *IAudioClient) SetEventHandle(handle uintptr) (err error) {
	err = acSetEventHandle(v, handle)
	return
}

func (v *IAudioClient) GetService(refIID *ole.GUID, obj interface{}) (err error) {
	err = acGetService(v, refIID, obj)
	return
}
