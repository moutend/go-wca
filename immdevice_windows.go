// +build windows

package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmdActivate(mmd *IMMDevice, refIID *ole.GUID, ctx uint32, prop, obj interface{}) (err error) {
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

func mmdOpenPropertyStore(mmd *IMMDevice, storageMode uint32, ps **IPropertyStore) (err error) {
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

func mmdGetId(mmd *IMMDevice, strId *string) (err error) {
	var strIdPtr uint32
	hr, _, _ := syscall.Syscall(
		mmd.VTable().GetId,
		2,
		uintptr(unsafe.Pointer(mmd)),
		uintptr(unsafe.Pointer(&strIdPtr)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	// According to the MSDN document, an endpoint ID string is a null-terminated wide-character string.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd370837(v=vs.85).aspx
	var us []uint16
	var i uint32
	var start = unsafe.Pointer(uintptr(strIdPtr))
	for {
		u := *(*uint16)(unsafe.Pointer(uintptr(start) + 2*uintptr(i)))
		if u == 0 {
			break
		}
		us = append(us, u)
		i++
	}
	*strId = syscall.UTF16ToString(us)
	ole.CoTaskMemFree(uintptr(strIdPtr))
	return
}

func mmdGetState(mmd *IMMDevice, state *uint32) (err error) {
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
