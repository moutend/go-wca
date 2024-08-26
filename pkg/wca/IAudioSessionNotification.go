package wca

import "github.com/go-ole/go-ole"

type IAudioSessionNotificationCallback struct {
	OnSessionCreated func(pNewSession *IAudioSessionControl) error
}

type IAudioSessionNotification struct {
	vTable   *IAudioSessionNotificationVtbl
	refCount int
	callback IAudioSessionNotificationCallback
}

type IAudioSessionNotificationVtbl struct {
	ole.IUnknownVtbl

	OnSessionCreated uintptr
}
