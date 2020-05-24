// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func mmdcGetCount(dc *IMMDeviceCollection, count *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdcItem(dc *IMMDeviceCollection, id uint32, mmd **IMMDevice) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
