//go:build windows

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
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetChannelCount,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(channelCount)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevSetMasterVolumeLevel(aev *IAudioEndpointVolume, levelDB float32, eventContextGUID *ole.GUID) (err error) {
	levelDBValue := math.Float32bits(levelDB)

	hr, _, _ := syscall.SyscallN(
		aev.VTable().SetMasterVolumeLevel,
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

	hr, _, _ := syscall.SyscallN(
		aev.VTable().SetMasterVolumeLevelScalar,
		uintptr(unsafe.Pointer(aev)),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetMasterVolumeLevel(aev *IAudioEndpointVolume, level *float32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetMasterVolumeLevel,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(level)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetMasterVolumeLevelScalar(aev *IAudioEndpointVolume, level *float32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetMasterVolumeLevelScalar,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(level)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevSetChannelVolumeLevel(aev *IAudioEndpointVolume, channel uint32, levelDB float32, eventContextGUID *ole.GUID) (err error) {
	levelDBValue := math.Float32bits(levelDB)

	hr, _, _ := syscall.SyscallN(
		aev.VTable().SetChannelVolumeLevel,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(levelDBValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevSetChannelVolumeLevelScalar(aev *IAudioEndpointVolume, channel uint32, level float32, eventContextGUID *ole.GUID) (err error) {
	levelValue := math.Float32bits(level)

	hr, _, _ := syscall.SyscallN(
		aev.VTable().SetChannelVolumeLevelScalar,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetChannelVolumeLevel(aev *IAudioEndpointVolume, channel uint32, levelDB *float32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetChannelVolumeLevel,
		uintptr(unsafe.Pointer(aev)),
		uintptr(channel),
		uintptr(unsafe.Pointer(levelDB)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetChannelVolumeLevelScalar(aev *IAudioEndpointVolume, channel uint32, level *float32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetChannelVolumeLevelScalar,
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
	hr, _, _ := syscall.SyscallN(
		aev.VTable().SetMute,
		uintptr(unsafe.Pointer(aev)),
		uintptr(muteValue),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetMute(aev *IAudioEndpointVolume, mute *bool) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetMute,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(mute)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetVolumeStepInfo(aev *IAudioEndpointVolume, step, stepCount *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetVolumeStepInfo,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(step)),
		uintptr(unsafe.Pointer(stepCount)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevVolumeStepUp(aev *IAudioEndpointVolume, eventContextGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().VolumeStepUp,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevVolumeStepDown(aev *IAudioEndpointVolume, eventContextGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().VolumeStepDown,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(eventContextGUID)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevQueryHardwareSupport(aev *IAudioEndpointVolume, hardwareSupportMask *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().QueryHardwareSupport,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(hardwareSupportMask)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func aevGetVolumeRange(aev *IAudioEndpointVolume, minDB, maxDB, incrementDB *float32) (err error) {
	hr, _, _ := syscall.SyscallN(
		aev.VTable().GetVolumeRange,
		uintptr(unsafe.Pointer(aev)),
		uintptr(unsafe.Pointer(minDB)),
		uintptr(unsafe.Pointer(maxDB)),
		uintptr(unsafe.Pointer(incrementDB)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
