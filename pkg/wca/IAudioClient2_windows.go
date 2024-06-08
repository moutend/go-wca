//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ac2IsOffloadCapable(ac2 *IAudioClient2, category uint32, isOffloadCapable *bool) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac2.VTable().IsOffloadCapable,
		uintptr(unsafe.Pointer(ac2)),
		uintptr(category),
		uintptr(unsafe.Pointer(isOffloadCapable)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ac2SetClientProperties(ac2 *IAudioClient2, properties *AudioClientProperties) (err error) {
	hr, _, _ := syscall.SyscallN(
		ac2.VTable().SetClientProperties,
		uintptr(unsafe.Pointer(ac2)),
		uintptr(unsafe.Pointer(properties)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ac2GetBufferSizeLimits(ac2 *IAudioClient2, wfx *WAVEFORMATEX, isEventDriven bool, minBufferDuration, maxBufferDuration *uint32) (err error) {
	var isEventDrivenValue uint32

	if isEventDriven {
		isEventDrivenValue = 1
	}
	hr, _, _ := syscall.SyscallN(
		ac2.VTable().GetBufferSizeLimits,
		uintptr(unsafe.Pointer(ac2)),
		uintptr(unsafe.Pointer(wfx)),
		uintptr(isEventDrivenValue),
		uintptr(unsafe.Pointer(minBufferDuration)),
		uintptr(unsafe.Pointer(maxBufferDuration)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
