package main

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMDevice struct {
	ole.IUnknown
}

type IMMDeviceVtbl struct {
	ole.IUnknownVtbl
	Activate          uintptr
	GetId             uintptr
	GetState          uintptr
	OpenPropertyStore uintptr
}

func (v *IMMDevice) VTable() *IMMDeviceVtbl {
	return (*IMMDeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMDevice) GetId(id **uint16) (err error) {
	err = getId(v, id)
	return
}
