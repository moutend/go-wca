// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func accGetBuffer(acc *IAudioCaptureClient, data **byte, framesToRead, flags *uint32, devicePosition, qpcPosition *uint64) (err error) {
	hr, _, _ := syscall.Syscall6(
		acc.VTable().GetBuffer,
		6,
		uintptr(unsafe.Pointer(acc)),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(framesToRead)),
		uintptr(unsafe.Pointer(flags)),
		uintptr(unsafe.Pointer(devicePosition)),
		uintptr(unsafe.Pointer(qpcPosition)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func accReleaseBuffer(acc *IAudioCaptureClient, framesRead uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		acc.VTable().ReleaseBuffer,
		2,
		uintptr(unsafe.Pointer(acc)),
		uintptr(framesRead),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func accGetNextPacketSize(acc *IAudioCaptureClient, framesInNextPacket *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		acc.VTable().GetNextPacketSize,
		2,
		uintptr(unsafe.Pointer(acc)),
		uintptr(unsafe.Pointer(framesInNextPacket)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
