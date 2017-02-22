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
func getMasterVolumeLevelScalar(aev *IAudioEndpointVolume, level *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetMasterVolumeLevelScalar,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(level)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
func setMasterVolumeLevelScalar(aev *IAudioEndpointVolume, level float32, eventContextGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().SetMasterVolumeLevelScalar,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(&level)),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func volumeStepUp(aev *IAudioEndpointVolume, eventContextGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().VolumeStepUp,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(eventContextGUID)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
