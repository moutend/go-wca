//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func arcGetBuffer(arc *IAudioRenderClient, requiredBufferSize uint32, data **byte) (err error) {
	hr, _, _ := syscall.SyscallN(
		arc.VTable().GetBuffer,
		uintptr(unsafe.Pointer(arc)),
		uintptr(requiredBufferSize),
		uintptr(unsafe.Pointer(data)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func arcReleaseBuffer(arc *IAudioRenderClient, writtenBufferSize, flag uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		arc.VTable().ReleaseBuffer,
		uintptr(unsafe.Pointer(arc)),
		uintptr(writtenBufferSize),
		uintptr(flag))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
