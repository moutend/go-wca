//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ascGetState(asc *IAudioSessionControl, retVal *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().GetState,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(retVal)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ascGetDisplayName(asc *IAudioSessionControl, retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.SyscallN(
		asc.VTable().GetDisplayName,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(&retValPtr)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	var us []uint16
	var i uint32
	var start = unsafe.Pointer(uintptr(retValPtr))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}
	*retVal = syscall.UTF16ToString(us)
	ole.CoTaskMemFree(uintptr(retValPtr))
	return
}

func ascSetDisplayName(asc *IAudioSessionControl, value *string, eventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().SetDisplayName,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(value)),
		uintptr(unsafe.Pointer(eventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ascGetIconPath(asc *IAudioSessionControl, retVal *string) (err error) {
	var retValPtr uint64
	hr, _, _ := syscall.SyscallN(
		asc.VTable().GetIconPath,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(&retValPtr)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	var us []uint16
	var i uint32
	var start = unsafe.Pointer(uintptr(retValPtr))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}
	*retVal = syscall.UTF16ToString(us)
	ole.CoTaskMemFree(uintptr(retValPtr))
	return
}

func ascSetIconPath(asc *IAudioSessionControl, value *string, eventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().SetIconPath,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(value)),
		uintptr(unsafe.Pointer(eventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ascGetGroupingParam(asc *IAudioSessionControl, retVal *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().GetGroupingParam,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(retVal)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ascSetGroupingParam(asc *IAudioSessionControl, override *ole.GUID, eventContext *ole.GUID) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().SetGroupingParam,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(override)),
		uintptr(unsafe.Pointer(eventContext)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ascRegisterAudioSessionNotification(asc *IAudioSessionControl, newNotifications *IAudioSessionEvents) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().RegisterAudioSessionNotification,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(newNotifications)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ascUnregisterAudioSessionNotification(asc *IAudioSessionControl, newNotifications *IAudioSessionEvents) (err error) {
	hr, _, _ := syscall.SyscallN(
		asc.VTable().UnregisterAudioSessionNotification,
		uintptr(unsafe.Pointer(asc)),
		uintptr(unsafe.Pointer(newNotifications)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
