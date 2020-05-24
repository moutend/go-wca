package wca

type IAudioSessionNotification struct {
	VTable *IAudioSessionEventsVtbl
}

type IAudioSessionNotificationVtbl struct {
	QueryInterface   uintptr
	AddRef           uintptr
	Release          uintptr
	OnSessionCreated uintptr
}
