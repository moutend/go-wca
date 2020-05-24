package wca

type IAudioVolumeDuckNotification struct {
	VTable *IAudioSessionEventsVtbl
}

type IAudioVolumeDuckNotificationVtbl struct {
	QueryInterface             uintptr
	AddRef                     uintptr
	Release                    uintptr
	OnVolumeDuckNotification   uintptr
	OnVolumeUnduckNotification uintptr
}
