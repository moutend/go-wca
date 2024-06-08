package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMDeviceCollection struct {
	ole.IUnknown
}

type IMMDeviceCollectionVtbl struct {
	ole.IUnknownVtbl
	GetCount uintptr
	Item     uintptr
}

func (v *IMMDeviceCollection) VTable() *IMMDeviceCollectionVtbl {
	return (*IMMDeviceCollectionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMDeviceCollection) GetCount(count *uint32) (err error) {
	err = mmdcGetCount(v, count)
	return
}

func (v *IMMDeviceCollection) Item(id uint32, mmd **IMMDevice) (err error) {
	err = mmdcItem(v, id, mmd)
	return
}
