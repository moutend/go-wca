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
