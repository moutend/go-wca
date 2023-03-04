//go:build windows
// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func amiGetPeakValue(ami *IAudioMeterInformation, peak *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		ami.VTable().GetPeakValue,
		2,
		uintptr(unsafe.Pointer(ami)),
		uintptr(unsafe.Pointer(peak)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return

}

func amiGetChannelsPeakValues(ami *IAudioMeterInformation, count uint32, peaks []float32) (err error) {
	hr, _, _ := syscall.Syscall(ami.VTable().GetChannelsPeakValues,
		3,
		uintptr(unsafe.Pointer(ami)),
		uintptr(count),
		uintptr(unsafe.Pointer(&peaks[0])))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func amiGetMeteringChannelCount(ami *IAudioMeterInformation, count *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ami.VTable().GetMeteringChannelCount,
		2,
		uintptr(unsafe.Pointer(ami)),
		uintptr(unsafe.Pointer(count)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func amiQueryHardwareSupport(ami *IAudioMeterInformation, response *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ami.VTable().GetMeteringChannelCount,
		2,
		uintptr(unsafe.Pointer(ami)),
		uintptr(unsafe.Pointer(response)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
