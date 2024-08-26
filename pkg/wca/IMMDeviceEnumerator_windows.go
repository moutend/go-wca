//go:build windows
// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmdeEnumAudioEndpoints(mmde *IMMDeviceEnumerator, eDataFlow, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	hr, _, _ := syscall.Syscall6(
		mmde.VTable().EnumAudioEndpoints,
		4,
		uintptr(unsafe.Pointer(mmde)),
		uintptr(eDataFlow), //uintptr(unsafe.Pointer(&eDataFlow)),
		uintptr(stateMask), //uintptr(unsafe.Pointer(&stateMask)),
		uintptr(unsafe.Pointer(dc)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdeGetDefaultAudioEndpoint(mmde *IMMDeviceEnumerator, eDataFlow, stateMask uint32, mmd **IMMDevice) (err error) {
	hr, _, _ := syscall.Syscall6(
		mmde.VTable().GetDefaultAudioEndpoint,
		4,
		uintptr(unsafe.Pointer(mmde)),
		uintptr(eDataFlow),
		uintptr(stateMask),
		uintptr(unsafe.Pointer(mmd)),
		0,
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdeGetDevice(mmde *IMMDeviceEnumerator, strId string, mmd **IMMDevice) (err error) {
	wstrId, err := syscall.UTF16FromString(strId)
	if err != nil {
		return
	}
	hr, _, _ := syscall.SyscallN(
		mmde.VTable().GetDevice,
		uintptr(unsafe.Pointer(mmde)),
		uintptr(unsafe.Pointer(unsafe.SliceData(wstrId))),
		uintptr(unsafe.Pointer(mmd)),
	)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdeRegisterEndpointNotificationCallback(mmde *IMMDeviceEnumerator, mmnc *IMMNotificationClient) (err error) {
	hr, _, _ := syscall.Syscall(
		mmde.VTable().RegisterEndpointNotificationCallback,
		2,
		uintptr(unsafe.Pointer(mmde)),
		uintptr(unsafe.Pointer(mmnc)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func mmdeUnregisterEndpointNotificationCallback(mmde *IMMDeviceEnumerator, mmnc *IMMNotificationClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
