package wca

import "github.com/go-ole/go-ole"

type IAudioSessionEventsCallback struct {
	OnDisplayNameChanged   func(newDisplayName string, eventContext *ole.GUID) error
	OnIconPathChanged      func(newIconPath string, eventContext *ole.GUID) error
	OnSimpleVolumeChanged  func(newVolume float32, mute bool, eventContext *ole.GUID) error
	OnChannelVolumeChanged func(channelCount int, newChannelVolumeArray []float32, changedChannel int, eventContext *ole.GUID) error
	OnGroupingParamChanged func(newGroupingParam, eventContext *ole.GUID) error
	OnStateChanged         func(newState AudioSessionState) error
	OnSessionDisconnected  func(disconnectReason AudioSessionDisconnectReason) error
}

type IAudioSessionEvents struct {
	vTable   *IAudioSessionEventsVtbl
	refCount int
	callback IAudioSessionEventsCallback
}

type IAudioSessionEventsVtbl struct {
	ole.IUnknownVtbl

	OnDisplayNameChanged   uintptr
	OnIconPathChanged      uintptr
	OnSimpleVolumeChanged  uintptr
	OnChannelVolumeChanged uintptr
	OnGroupingParamChanged uintptr
	OnStateChanged         uintptr
	OnSessionDisconnected  uintptr
}
