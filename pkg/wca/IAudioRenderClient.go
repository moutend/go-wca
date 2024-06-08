package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioRenderClient struct {
	ole.IUnknown
}

type IAudioRenderClientVtbl struct {
	ole.IUnknownVtbl
	GetBuffer     uintptr
	ReleaseBuffer uintptr
}

func (v *IAudioRenderClient) VTable() *IAudioRenderClientVtbl {
	return (*IAudioRenderClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioRenderClient) GetBuffer(requiredFrameSize uint32, data **byte) (err error) {
	err = arcGetBuffer(v, requiredFrameSize, data)
	return
}

func (v *IAudioRenderClient) ReleaseBuffer(writtenFrameSize, flag uint32) (err error) {
	err = arcReleaseBuffer(v, writtenFrameSize, flag)
	return
}
