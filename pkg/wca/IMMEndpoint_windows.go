//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmeGetDataFlow(mme *IMMEndpoint, eDataFlow *uint32) (err error) {
	hr, _, _ := syscall.SyscallN(
		mme.VTable().GetDataFlow,
		uintptr(unsafe.Pointer(mme)),
		uintptr(unsafe.Pointer(eDataFlow)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
