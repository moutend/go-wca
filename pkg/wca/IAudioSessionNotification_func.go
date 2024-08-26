//go:build !windows
// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func asnQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	return ole.E_NOTIMPL
}

func asnAddRef(this uintptr) int64 {
	return 0
}

func asnRelease(this uintptr) int64 {
	return 0
}

func asnOnSessionCreated(this uintptr, pNewSession uintptr) int64 {
	return ole.E_NOTIMPL
}

func NewIAudioSessionNotification(callback IAudioSessionNotificationCallback) *IAudioSessionNotification {
	asn := &IAudioSessionNotification{}

	return asn
}
