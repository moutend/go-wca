// +build !windows

package main

import (
	"github.com/go-ole/go-ole"
)

func mmdeEnumAudioEndpoints(de *IMMDeviceEnumerator, eDataFlow, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeGetDefaultAudioEndpoint() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeGetDevice() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeRegisterEndpointNotificationCallback() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeUnregisterEndpointNotificationCallback() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
