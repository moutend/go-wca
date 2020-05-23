// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func mmncQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	*ppInterface = 0

	if ole.IsEqualGUID(riid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(riid, IID_IMMNotificationClient) {
		mmncAddRef(this)
		*ppInterface = this

		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func mmncAddRef(this uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	mmnc.refCount += 1

	return int64(mmnc.refCount)
}

func mmncRelease(this uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	mmnc.refCount -= 1

	return int64(mmnc.refCount)
}

func mmncOnDefaultDeviceChanged(this uintptr, flow, role uint64, pwstrDeviceId uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	if mmnc.callback.OnDefaultDeviceChanged == nil {
		return ole.S_OK
	}

	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	device := LPCWSTRToString(pwstrDeviceId, 1024)

	err := mmnc.callback.OnDefaultDeviceChanged(EDataFlow(flow), ERole(role), device)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func mmncOnDeviceAdded(this uintptr, pwstrDeviceId uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	if mmnc.callback.OnDeviceAdded == nil {
		return ole.S_OK
	}
	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	device := LPCWSTRToString(pwstrDeviceId, 1024)

	err := mmnc.callback.OnDeviceAdded(device)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func mmncOnDeviceRemoved(this uintptr, pwstrDeviceId uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	if mmnc.callback.OnDeviceRemoved == nil {
		return ole.S_OK
	}

	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	device := LPCWSTRToString(pwstrDeviceId, 1024)

	err := mmnc.callback.OnDeviceRemoved(device)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func mmncOnDeviceStateChanged(this uintptr, pwstrDeviceId uintptr, dwNewState uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	if mmnc.callback.OnDeviceStateChanged == nil {
		return ole.S_OK
	}

	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	device := LPCWSTRToString(pwstrDeviceId, 1024)

	err := mmnc.callback.OnDeviceStateChanged(device, 0)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func mmncOnPropertyValueChanged(this uintptr, pwstrDeviceId uintptr, key uintptr) int64 {
	mmnc := (*IMMNotificationClient)(unsafe.Pointer(this))

	if mmnc.callback.OnPropertyValueChanged == nil {
		return ole.S_OK
	}

	// device := syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(pwstrDeviceId)))
	device := LPCWSTRToString(pwstrDeviceId, 1024)

	err := mmnc.callback.OnPropertyValueChanged(device, 0)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

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

func NewIMMNotificationClient(callback IMMNotificationClientCallback) *IMMNotificationClient {
	vTable := &IMMNotificationClientVtbl{}

	// IUnknown methods
	vTable.QueryInterface = syscall.NewCallback(mmncQueryInterface)
	vTable.AddRef = syscall.NewCallback(mmncAddRef)
	vTable.Release = syscall.NewCallback(mmncRelease)

	// IMMNotificationClient methods
	vTable.OnDeviceStateChanged = syscall.NewCallback(mmncOnDeviceStateChanged)
	vTable.OnDeviceAdded = syscall.NewCallback(mmncOnDeviceAdded)
	vTable.OnDeviceRemoved = syscall.NewCallback(mmncOnDeviceRemoved)
	vTable.OnDefaultDeviceChanged = syscall.NewCallback(mmncOnDefaultDeviceChanged)
	vTable.OnPropertyValueChanged = syscall.NewCallback(mmncOnPropertyValueChanged)

	mmnc := &IMMNotificationClient{}

	mmnc.vTable = vTable
	mmnc.callback = callback

	return mmnc
}
