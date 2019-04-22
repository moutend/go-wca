// +build !windows

package wca

import "github.com/go-ole/go-ole"

func ascGetState(asc *IAudioSessionControl, retVal *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascGetDisplayName(asc *IAudioSessionControl, retVal *string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascSetDisplayName(asc *IAudioSessionControl, value *string, eventContext *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascGetIconPath(asc *IAudioSessionControl, retVal *string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascSetIconPath(asc *IAudioSessionControl, value *string, eventContext *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascGetGroupingParam(asc *IAudioSessionControl, retVal *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascSetGroupingParam(asc *IAudioSessionControl, override *ole.GUID, eventContext *ole.GUID) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascRegisterAudioSessionNotification(asc *IAudioSessionControl, newNotifications *IAudioSessionEvents) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func ascUnregisterAudioSessionNotification(asc *IAudioSessionControl, newNotifications *IAudioSessionEvents) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
