package main

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
	err = getCount(v, count)
	return
}
