// +build !windows

package wca

import "github.com/go-ole/go-ole"

func savSetMasterVolume(sav *ISimpleAudioVolume, level float32, eventContext *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func savGetMasterVolume(sav *ISimpleAudioVolume, level *float32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func savSetMute(sav *ISimpleAudioVolume, mute bool, eventContext *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func savGetMute(sav *ISimpleAudioVolume, mute *bool) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
