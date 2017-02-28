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

func (v *IPropertyStore) GetCount(count *uint32) (err error) {
	err = psGetCount(v, count)
	return
}

func (v *IPropertyStore) GetAt(index uint32, pk *PropertyKey) (err error) {
	err = psGetAt(v, index, pk)
	return
}

func (v *IPropertyStore) GetValue(key *PropertyKey, pv *PROPVARIANT) (err error) {
	err = psGetValue(v, key, pv)
	return
}

func (v *IPropertyStore) SetValue() (err error) {
	return
}

func (v *IPropertyStore) Commit() (err error) {
	return
}
