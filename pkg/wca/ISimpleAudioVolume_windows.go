// +build windows

package wca

import (
	"math"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func savSetMasterVolume(sav *ISimpleAudioVolume, level float32, eventContext *ole.GUID) (err error) {
	levelValue := math.Float32bits(level)

	hr, _, _ := syscall.Syscall(
		sav.VTable().SetMasterVolume,
		3,
		uintptr(unsafe.Pointer(sav)),
		uintptr(levelValue),
		uintptr(unsafe.Pointer(eventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func savGetMasterVolume(sav *ISimpleAudioVolume, level *float32) (err error) {
	hr, _, _ := syscall.Syscall(
		sav.VTable().GetMasterVolume,
		2,
		uintptr(unsafe.Pointer(sav)),
		uintptr(unsafe.Pointer(level)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func savSetMute(sav *ISimpleAudioVolume, mute bool, eventContext *ole.GUID) (err error) {
	var muteValue uint32

	if mute {
		muteValue = 1
	}
	hr, _, _ := syscall.Syscall(
		sav.VTable().SetMute,
		3,
		uintptr(unsafe.Pointer(sav)),
		uintptr(muteValue),
		uintptr(unsafe.Pointer(eventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func savGetMute(sav *ISimpleAudioVolume, mute *bool) (err error) {
	hr, _, _ := syscall.Syscall(
		sav.VTable().GetMute,
		2,
		uintptr(unsafe.Pointer(sav)),
		uintptr(unsafe.Pointer(mute)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
