// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func pvString(v int64) (s string) {
	var us []uint16
	var i uint32
	var start = unsafe.Pointer(uintptr(v))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}
	s = syscall.UTF16ToString(us)
	ole.CoTaskMemFree(uintptr(v))
	return
}
