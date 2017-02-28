// +build !windows

package main

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func psGetCount(ps *IPropertyStore, count *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psGetAt(ps *IPropertyStore, index uint32, pk *PropertyKey) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psGetValue(ps *IPropertyStore, key *PropertyKey, pv *PROPVARIANT) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psSetValue() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psCommit() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
