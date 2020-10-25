package wca

type IAudioSessionNotification struct {
	VTable *IAudioSessionNotificationVtbl
}

type IAudioSessionNotificationVtbl struct {
	QueryInterface   uintptr
	AddRef           uintptr
	Release          uintptr
	OnSessionCreated uintptr
}
