// +build windows

package main

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func activate(mmd *IMMDevice, refIID *ole.GUID, ctx uint32, prop, obj interface{}) (err error) {
	objValue := reflect.ValueOf(obj).Elem()
	hr, _, _ := syscall.Syscall6(
		mmd.VTable().Activate,
		5,
		uintptr(unsafe.Pointer(mmd)),
		uintptr(unsafe.Pointer(refIID)),
		uintptr(unsafe.Pointer(&ctx)),
		0,
		objValue.Addr().Pointer(),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func openPropertyStore(mmd *IMMDevice, storageMode uint32, ps **IPropertyStore) (err error) {
	hr, _, _ := syscall.Syscall(
		mmd.VTable().OpenPropertyStore,
		3,
		uintptr(unsafe.Pointer(mmd)),
		uintptr(storageMode),
		uintptr(unsafe.Pointer(ps)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getId(mmd *IMMDevice, strId *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		mmd.VTable().GetId,
		2,
		uintptr(unsafe.Pointer(mmd)),
		uintptr(unsafe.Pointer(strId)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func getState(mmd *IMMDevice, state *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		mmd.VTable().GetState,
		2,
		uintptr(unsafe.Pointer(mmd)),
		uintptr(unsafe.Pointer(state)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
