package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMNotificationClient struct {
	ole.IUnknown
}

type IMMNotificationClientVtbl struct {
	ole.IUnknownVtbl
	OnDeviceStateChanged   uintptr
	OnDeviceAdded          uintptr
	OnDeviceRemoved        uintptr
	OnDefaultDeviceChanged uintptr
	OnPropertyValueChanged uintptr
}

func (v *IMMNotificationClient) VTable() *IMMNotificationClientVtbl {
	return (*IMMNotificationClientVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMNotificationClient) OnDeviceStateChanged() (err error) {
	return
}

func (v *IMMNotificationClient) OnDeviceAdded() (err error) {
	return
}

func (v *IMMNotificationClient) OnDeviceRemoved() (err error) {
	return
}

func (v *IMMNotificationClient) OnDefaultDeviceChanged() (err error) {
	return
}

func (v *IMMNotificationClient) OnPropertyValueChanged() (err error) {
	return
}
