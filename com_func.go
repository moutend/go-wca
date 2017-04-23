// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func CreateEventExA(securityAttributes, name, flag, desiredAccess uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func CloseHandle(hObject uintptr) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func CoCreateInstance(clsid *ole.GUID, punk uintptr, clsctx uint32, iid *ole.GUID, obj interface{}) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func WaitForSingleObject(handle uintptr, milliseconds uint32) (dword uint32) {
	return
}
