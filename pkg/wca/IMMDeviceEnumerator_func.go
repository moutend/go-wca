// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func mmdeEnumAudioEndpoints(mmde *IMMDeviceEnumerator, eDataFlow, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeGetDefaultAudioEndpoint(mmde *IMMDeviceEnumerator, eDataFlow, stateMask uint32, mmd **IMMDevice) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeGetDevice() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeRegisterEndpointNotificationCallback(mmde *IMMDeviceEnumerator, mmnc *IMMNotificationClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdeUnregisterEndpointNotificationCallback(mmde *IMMDeviceEnumerator, mmnc *IMMNotificationClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
