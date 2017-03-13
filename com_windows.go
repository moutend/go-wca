// +build windows
package wca

import (
	"github.com/go-ole/go-ole"
	"syscall"
)

var (
	modkernel32, _       = syscall.LoadDLL("kernel32.dll")
	procCreateEventEx, _ = modkernel32.FindProc("CreateEventEx")
)

func CreateEventEx(securityAttributes, name, flag, desiredAccess uint32) (err error) {
	hr, _, _ := procCreateEventEx.Call(
		uintptr(securityAttributes),
		uintptr(name),
		uintptr(flag),
		uintptr(desiredAccess))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
