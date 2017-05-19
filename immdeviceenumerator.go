package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMDeviceEnumerator struct {
	ole.IUnknown
}

type IMMDeviceEnumeratorVtbl struct {
	ole.IUnknownVtbl
	EnumAudioEndpoints                     uintptr
	GetDefaultAudioEndpoint                uintptr
	GetDevice                              uintptr
	RegisterEndpointNotificationCallback   uintptr
	UnregisterEndpointNotificationCallback uintptr
}

func (v *IMMDeviceEnumerator) VTable() *IMMDeviceEnumeratorVtbl {
	return (*IMMDeviceEnumeratorVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMDeviceEnumerator) EnumAudioEndpoints(eDataFlow, stateMask uint32, dc **IMMDeviceCollection) (err error) {
	err = mmdeEnumAudioEndpoints(v, eDataFlow, stateMask, dc)
	return
}

func (v *IMMDeviceEnumerator) GetDefaultAudioEndpoint(eDataFlow, stateMask uint32, mmd **IMMDevice) (err error) {
	err = mmdeGetDefaultAudioEndpoint(v, eDataFlow, stateMask, mmd)
	return
}

func (v *IMMDeviceEnumerator) GetDevice() (err error) {
	err = mmdeGetDevice()
	return
}

func (v *IMMDeviceEnumerator) RegisterEndpointNotificationCallback(mmnc *IMMNotificationClient) (err error) {
	err = mmdeRegisterEndpointNotificationCallback(v, mmnc)
	return
}

func (v *IMMDeviceEnumerator) UnregisterEndpointNotificationCallback(mmnc *IMMNotificationClient) (err error) {
	err = mmdeUnregisterEndpointNotificationCallback(v, mmnc)
	return
}
