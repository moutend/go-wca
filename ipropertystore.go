package main

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IPropertyStore struct {
	ole.IUnknown
}

type IPropertyStoreVtbl struct {
	ole.IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	GetValue uintptr
	SetValue uintptr
	Commit   uintptr
}

func (v *IPropertyStore) VTable() *IPropertyStoreVtbl {
	return (*IPropertyStoreVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IPropertyStore) GetValue(key *PropertyKey, pv *PROPVARIANT) (err error) {
	err = getValue(v, key, pv)
	return
}
