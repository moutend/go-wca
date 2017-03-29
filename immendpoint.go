package wca

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type IMMEndpoint struct {
	ole.IUnknown
}

type IMMEndpointVtbl struct {
	ole.IUnknownVtbl
	GetDataFlow uintptr
}

func (v *IMMEndpoint) VTable() *IMMEndpointVtbl {
	return (*IMMEndpointVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IMMEndpoint) GetDataFlow(eDataFlow *uint32) (err error) {
	err = mmeGetDataFlow(v, eDataFlow)
	return
}
