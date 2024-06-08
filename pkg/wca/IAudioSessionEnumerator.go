package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioSessionEnumerator struct {
	ole.IUnknown
}

type IAudioSessionEnumeratorVtbl struct {
	ole.IUnknownVtbl
	GetCount   uintptr
	GetSession uintptr
}

func (v *IAudioSessionEnumerator) VTable() *IAudioSessionEnumeratorVtbl {
	return (*IAudioSessionEnumeratorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioSessionEnumerator) GetCount(sessionCount *int) (err error) {
	err = aseGetCount(v, sessionCount)
	return
}

func (v *IAudioSessionEnumerator) GetSession(sessionCount int, session **IAudioSessionControl) (err error) {
	err = aseGetSession(v, sessionCount, session)
	return
}
