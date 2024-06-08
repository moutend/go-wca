package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IAudioCaptureClient struct {
	ole.IUnknown
}

type IAudioCaptureClientVtbl struct {
	ole.IUnknownVtbl
	GetBuffer         uintptr
	ReleaseBuffer     uintptr
	GetNextPacketSize uintptr
}

func (v *IAudioCaptureClient) VTable() *IAudioCaptureClientVtbl {
	return (*IAudioCaptureClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IAudioCaptureClient) GetBuffer(data **byte, framesToRead, flags *uint32, devicePosition, qpcPosition *uint64) (err error) {
	err = accGetBuffer(v, data, framesToRead, flags, devicePosition, qpcPosition)
	return
}

func (v *IAudioCaptureClient) ReleaseBuffer(framesRead uint32) (err error) {
	err = accReleaseBuffer(v, framesRead)
	return
}

func (v *IAudioCaptureClient) GetNextPacketSize(framesInNextPacket *uint32) (err error) {
	err = accGetNextPacketSize(v, framesInNextPacket)
	return
}
