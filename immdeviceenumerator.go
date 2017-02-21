package main

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
	err = enumAudioEndpoints(v, eDataFlow, stateMask, dc)
	return
}
