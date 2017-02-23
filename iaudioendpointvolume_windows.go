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
		0) //uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getVolumeStepInfo(aev *IAudioEndpointVolume, step, stepCount *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetVolumeStepInfo,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(step)),
		uintptr(unsafe.Pointer(stepCount)))
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

func volumeStepDown(aev *IAudioEndpointVolume, eventContextGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().VolumeStepDown,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(eventContextGUID)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getVolumeRange(aev *IAudioEndpointVolume, minDB, maxDB, incrementDB *float32) (err error) {
	hr, _, _ := syscall.Syscall6(
		aev.VTable().GetVolumeRange,
		4,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(minDB)),
		uintptr(unsafe.Pointer(maxDB)),
		uintptr(unsafe.Pointer(incrementDB)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func queryHardwareSupport(aev *IAudioEndpointVolume, hardwareSupportMask *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().QueryHardwareSupport,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(hardwareSupportMask)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
