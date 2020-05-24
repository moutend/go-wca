package wca

import (
	"github.com/go-ole/go-ole"
)

type IMMNotificationClientCallback struct {
	OnDefaultDeviceChanged func(flow EDataFlow, role ERole, pwstrDeviceId string) error
	OnDeviceAdded          func(pwstrDeviceId string) error
	OnDeviceRemoved        func(pwstrDeviceId string) error
	OnDeviceStateChanged   func(pwstrDeviceId string, dwNewState uint64) error
	OnPropertyValueChanged func(pwstrDeviceId string, key uint64) error
}

type IMMNotificationClient struct {
	vTable   *IMMNotificationClientVtbl
	refCount int
	callback IMMNotificationClientCallback
}

type IMMNotificationClientVtbl struct {
	ole.IUnknownVtbl

	OnDeviceStateChanged   uintptr
	OnDeviceAdded          uintptr
	OnDeviceRemoved        uintptr
	OnDefaultDeviceChanged uintptr
	OnPropertyValueChanged uintptr
}
