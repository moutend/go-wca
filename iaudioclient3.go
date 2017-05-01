package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IAudioClient3 struct corresponds to IAudioClient3 interface.
// See MSDN's documentation:
// https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx
type IAudioClient3 struct {
	ole.IUnknown
	IAudioClient
	IAudioClient2
}

type IAudioClient3Vtbl struct {
	ole.IUnknownVtbl
	IAudioClientVtbl
	IAudioClient2Vtbl
	GetSharedModeEnginePeriod        uintptr
	GetCurrentSharedModeEnginePeriod uintptr
	InitializeSharedAudioStream      uintptr
}

func (v *IAudioClient3) VTable() *IAudioClient3Vtbl {
	return (*IAudioClient3Vtbl)(unsafe.Pointer(v.RawVTable))
}
