//go:build windows
// +build windows

package wca

import (
	"syscall"
	"unsafe"
)

func LPCWSTRToString(lpcwstr uintptr, maxChars int) string {
	if lpcwstr == 0 || maxChars == 0 {
		return ""
	}

	us := []uint16{}

	for i := 0; i < maxChars; i += 2 {
		u := *(*uint16)(unsafe.Pointer(lpcwstr + uintptr(i)))

		if u == 0 {
			break
		}

		us = append(us, u)
	}

	return syscall.UTF16ToString(us)
}
