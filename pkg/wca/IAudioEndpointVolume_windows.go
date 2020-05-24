// +build windows

package wca

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func aevRegisterControlChangeNotify() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func aevUnregisterControlChangeNotify() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func aevGetChannelCount(aev *IAudioEndpointVolume, channelCount *uint32) (err error) {
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

func aevSetMasterVolumeLevel(aev *IAudioEndpointVolume, levelDB float32, eventContextGUID *ole.GUID) (err error) {
	levelDBValue := math.Float32bits(levelDB)

	hr, _, _ := syscall.Syscall(
		aev.VTable().SetMasterVolumeLevel,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(levelDBValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevSetMasterVolumeLevelScalar(aev *IAudioEndpointVolume, level float32, eventContextGUID *ole.GUID) (err error) {
	levelValue := math.Float32bits(level)

	hr, _, _ := syscall.Syscall(
		aev.VTable().SetMasterVolumeLevelScalar,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetMasterVolumeLevel(aev *IAudioEndpointVolume, level *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetMasterVolumeLevel,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(level)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetMasterVolumeLevelScalar(aev *IAudioEndpointVolume, level *float32) (err error) {
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

func aevSetChannelVolumeLevel(aev *IAudioEndpointVolume, channel uint32, levelDB float32, eventContextGUID *ole.GUID) (err error) {
	levelDBValue := math.Float32bits(levelDB)

	hr, _, _ := syscall.Syscall6(
		aev.VTable().SetChannelVolumeLevel,
		4,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(levelDBValue),
		uintptr(unsafe.Pointer(eventContextGUID)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevSetChannelVolumeLevelScalar(aev *IAudioEndpointVolume, channel uint32, level float32, eventContextGUID *ole.GUID) (err error) {
	levelValue := math.Float32bits(level)

	hr, _, _ := syscall.Syscall6(
		aev.VTable().SetChannelVolumeLevelScalar,
		4,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(eventContextGUID)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetChannelVolumeLevel(aev *IAudioEndpointVolume, channel uint32, levelDB *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetChannelVolumeLevel,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(unsafe.Pointer(levelDB)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetChannelVolumeLevelScalar(aev *IAudioEndpointVolume, channel uint32, level *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetChannelVolumeLevelScalar,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(unsafe.Pointer(level)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevSetMute(aev *IAudioEndpointVolume, mute bool, eventContextGUID *ole.GUID) (err error) {
	var muteValue uint32

	if mute {
		muteValue = 1
	}
	hr, _, _ := syscall.Syscall(
		aev.VTable().SetMute,
		3,
		uintptr(unsafe.Pointer(aev)),
		uintptr(muteValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetMute(aev *IAudioEndpointVolume, mute *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		aev.VTable().GetMute,
		2,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(mute)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetVolumeStepInfo(aev *IAudioEndpointVolume, step, stepCount *uint32) (err error) {
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

func aevVolumeStepUp(aev *IAudioEndpointVolume, eventContextGUID *ole.GUID) (err error) {
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

func aevVolumeStepDown(aev *IAudioEndpointVolume, eventContextGUID *ole.GUID) (err error) {
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

func aevQueryHardwareSupport(aev *IAudioEndpointVolume, hardwareSupportMask *uint32) (err error) {
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

func aevGetVolumeRange(aev *IAudioEndpointVolume, minDB, maxDB, incrementDB *float32) (err error) {
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
