// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmeGetDataFlow(mme *IMMEndpoint, eDataFlow *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		mme.VTable().GetDataFlow,
		2,
		uintptr(unsafe.Pointer(mme)),
		uintptr(unsafe.Pointer(eDataFlow)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
