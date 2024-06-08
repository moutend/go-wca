package wca

type IAudioSessionEvents struct {
	VTable *IAudioSessionEventsVtbl
}

type IAudioSessionEventsVtbl struct {
	QueryInterface         uintptr
	AddRef                 uintptr
	Release                uintptr
	OnDisplayNameChanged   uintptr
	OnIconPathChanged      uintptr
	OnSimpleVolumeChanged  uintptr
	OnChannelVolumeChanged uintptr
	OnGroupingParamChanged uintptr
	OnStateChanged         uintptr
	OnSessionDisconnected  uintptr
}
