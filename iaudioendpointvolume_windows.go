// +build windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func getChannelCount(aev *IAudioEndpointVolume, channelCount *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetChannelCount,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(channelCount)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
