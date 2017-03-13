// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func CreateEventEx(securityAttributes, name, flag, desiredAccess uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
