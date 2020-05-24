// +build windows
package wca

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modkernel32, _ = syscall.LoadDLL("kernel32.dll")
	modole32, _    = syscall.LoadDLL("ole32.dll")

	procCreateEventExA, _      = modkernel32.FindProc("CreateEventExA")
	procCloseHandle, _         = modkernel32.FindProc("CloseHandle")
	procCoCreateInstance, _    = modole32.FindProc("CoCreateInstance")
	procWaitForSingleObject, _ = modkernel32.FindProc("WaitForSingleObject")
)

func CreateEventExA(securityAttributes, name uintptr, flag, desiredAccess uint32) (handle uintptr) {
	handle, _, _ = procCreateEventExA.Call(
		securityAttributes,
		name,
		uintptr(flag),
		uintptr(desiredAccess))
	return
}

func CloseHandle(hObject uintptr) (err error) {
	hr, _, _ := procCloseHandle.Call(uintptr(hObject))
	if hr == 0 {
		err = fmt.Errorf("unexpected error: call GetLastError to get details")
	}
	return
}

func CoCreateInstance(clsid *ole.GUID, punk uintptr, clsctx uint32, iid *ole.GUID, obj interface{}) (err error) {
	objValue := reflect.ValueOf(obj).Elem()
	hr, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsid)),
		punk,
		uintptr(clsctx),
		uintptr(unsafe.Pointer(iid)),
		objValue.Addr().Pointer())
	if hr != ole.S_OK {
		err = ole.NewError(hr)
	}
	return
}

func WaitForSingleObject(handle uintptr, milliseconds uint32) (dword uint32) {
	hr, _, _ := procWaitForSingleObject.Call(
		uintptr(handle),
		uintptr(milliseconds))
	dword = uint32(hr)
	return
}
