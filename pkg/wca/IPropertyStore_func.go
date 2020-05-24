// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func psGetCount(ps *IPropertyStore, count *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psGetAt(ps *IPropertyStore, iProp uint32, pkey *PROPERTYKEY) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psGetValue(ps *IPropertyStore, key *PROPERTYKEY, pv *PROPVARIANT) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psSetValue() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func psCommit() (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
