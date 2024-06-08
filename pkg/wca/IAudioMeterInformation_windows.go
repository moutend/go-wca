//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func amiGetPeakValue(ami *IAudioMeterInformation, peak *float32) (err error) {
	hr, _, _ := syscall.SyscallN(
		ami.VTable().GetPeakValue,
		uintptr(unsafe.Pointer(ami)),
		uintptr(unsafe.Pointer(peak)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return

}

func amiGetChannelsPeakValues(ami *IAudioMeterInformation, count uint32, peaks []float32) (err error) {
	hr, _, _ := syscall.SyscallN(ami.VTable().GetChannelsPeakValues,
		uintptr(unsafe.Pointer(ami)),
		uintptr(count),
		uintptr(unsafe.Pointer(&peaks[0])))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func amiGetMeteringChannelCount(ami *IAudioMeterInformation, count *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		ami.VTable().GetMeteringChannelCount,
		uintptr(unsafe.Pointer(ami)),
		uintptr(unsafe.Pointer(count)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func amiQueryHardwareSupport(ami *IAudioMeterInformation, response *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		ami.VTable().GetMeteringChannelCount,
		uintptr(unsafe.Pointer(ami)),
		uintptr(unsafe.Pointer(response)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
