package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IAudioClient2 struct corresponds to IAudioClient2 interface.
// See MSDN's documentation:
// https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx
type IAudioClient2 struct {
	ole.IUnknown
	IAudioClient
}

type IAudioClient2Vtbl struct {
	ole.IUnknownVtbl
	IAudioClientVtbl
	IsOffloadCapable    uintptr
	SetClientProperties uintptr
	GetBufferSizeLimits uintptr
}

func (v *IAudioClient2) VTable() *IAudioClient2Vtbl {
	return (*IAudioClient2Vtbl)(unsafe.Pointer(v.RawVTable))
}
