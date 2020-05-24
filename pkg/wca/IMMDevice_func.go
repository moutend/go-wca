// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func mmdActivate(mmd *IMMDevice, refIID *ole.GUID, ctx uint32, prop, obj interface{}) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdOpenPropertyStore(mmd *IMMDevice, storageMode uint32, ps **IPropertyStore) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdGetId(mmd *IMMDevice, strId *string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func mmdGetState(mmd *IMMDevice, state *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
