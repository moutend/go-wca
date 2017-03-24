// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func arcGetBuffer(arc *IAudioRenderClient, requiredBufferSize uint32, data **byte) (err error) {
	hr, _, _ := syscall.Syscall(
		arc.VTable().GetBuffer,
		3,
		uintptr(unsafe.Pointer(arc)),
		uintptr(requiredBufferSize),
		uintptr(unsafe.Pointer(data)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func arcReleaseBuffer(arc *IAudioRenderClient, writtenBufferSize, flag uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		arc.VTable().ReleaseBuffer,
		3,
		uintptr(unsafe.Pointer(arc)),
		uintptr(writtenBufferSize),
		uintptr(flag))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
