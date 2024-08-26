//go:build !windows
// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func aseQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	return ole.E_NOTIMPL
}

func aseAddRef(this uintptr) int64 {
	return 0
}

func aseRelease(this uintptr) int64 {
	return 0
}

func aseOnDisplayNameChanged(this, lpcwstrNewDisplayName, lpcguidEventContext uintptr) int64 {
	return ole.E_NOTIMPL
}

func aseOnIconPathChanged(this, lpcwstrNewIconPath, lpcguidEventContext uintptr) int64 {
	return ole.E_NOTIMPL
}

func aseOnSimpleVolumeChanged(this uintptr, newVolume float32, muteValue uint32, lpcguidEventContext uintptr) int64 {
	return ole.E_NOTIMPL
}

func aseOnChannelVolumeChanged(this uintptr, channelCount uint32, newChannelVolumeArray []float32, changedChannel uint32, lpcguidEventContext uintptr) int64 {
	return ole.E_NOTIMPL
}

func aseOnGroupingParamChanged(this, lpcguidNewGroupingParam, lpcguidEventContext uintptr) int64 {
	return ole.E_NOTIMPL
}

func aseOnStateChanged(this uintptr, newState int64) int64 {
	return ole.E_NOTIMPL
}

func aseOnSessionDisconnected(this uintptr, disconnectReason int64) int64 {
	return ole.E_NOTIMPL
}

func NewIAudioSessionEvents(callback IAudioSessionEventsCallback) *IAudioSessionEvents {
	ase := &IAudioSessionEvents{}

	return ase
}
