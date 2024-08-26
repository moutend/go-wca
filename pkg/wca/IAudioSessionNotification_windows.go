//go:build windows
// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func asnQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	*ppInterface = 0

	if ole.IsEqualGUID(riid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(riid, IID_IAudioSessionNotification) {
		asnAddRef(this)
		*ppInterface = this

		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func asnAddRef(this uintptr) int64 {
	asn := (*IAudioSessionNotification)(unsafe.Pointer(this))

	asn.refCount += 1

	return int64(asn.refCount)
}

func asnRelease(this uintptr) int64 {
	asn := (*IAudioSessionNotification)(unsafe.Pointer(this))

	asn.refCount -= 1

	return int64(asn.refCount)
}

func asnOnSessionCreated(this uintptr, pNewSession uintptr) int64 {
	asn := (*IAudioSessionNotification)(unsafe.Pointer(this))

	if asn.callback.OnSessionCreated == nil {
		return ole.S_OK
	}

	session := (*IAudioSessionControl)(unsafe.Pointer(pNewSession))

	err := asn.callback.OnSessionCreated(session)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func NewIAudioSessionNotification(callback IAudioSessionNotificationCallback) *IAudioSessionNotification {
	vTable := &IAudioSessionNotificationVtbl{}

	// IUnknown methods
	vTable.QueryInterface = syscall.NewCallback(asnQueryInterface)
	vTable.AddRef = syscall.NewCallback(asnAddRef)
	vTable.Release = syscall.NewCallback(asnRelease)

	// IAudioSessionNotification methods
	vTable.OnSessionCreated = syscall.NewCallback(asnOnSessionCreated)

	asn := &IAudioSessionNotification{}

	asn.vTable = vTable
	asn.callback = callback

	return asn
}
