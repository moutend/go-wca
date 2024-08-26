//go:build !windows
// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func mmncQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	return ole.E_NOTIMPL
}

func mmncAddRef(this uintptr) int64 {
	return 0
}

func mmncRelease(this uintptr) int64 {
	return 0
}

func mmncOnDefaultDeviceChanged(this uintptr, flow, role uint64, pwstrDeviceId uintptr) int64 {
	return ole.E_NOTIMPL
}

func mmncOnDeviceAdded(this uintptr, pwstrDeviceId uintptr) int64 {
	return ole.E_NOTIMPL
}

func mmncOnDeviceRemoved(this uintptr, pwstrDeviceId uintptr) int64 {
	return ole.E_NOTIMPL
}

func mmncOnDeviceStateChanged(this uintptr, pwstrDeviceId uintptr, dwNewState uint32) int64 {
	return ole.E_NOTIMPL
}

func mmncOnPropertyValueChanged(this uintptr, pwstrDeviceId uintptr, key uintptr) int64 {
	return ole.E_NOTIMPL
}

func NewIMMNotificationClient(callback IMMNotificationClientCallback) *IMMNotificationClient {
	mmnc := &IMMNotificationClient{}

	return mmnc
}
