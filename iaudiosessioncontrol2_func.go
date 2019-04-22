// +build !windows

package wca

import "github.com/go-ole/go-ole"

func asc2GetSessionIdentifier(asc2 *IAudioSessionControl2, retVal *string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asc2GetSessionInstanceIdentifier(asc2 *IAudioSessionControl2, retVal *string) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asc2GetProcessId(asc2 *IAudioSessionControl2, retVal *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asc2IsSystemSoundsSession(asc2 *IAudioSessionControl2) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asc2SetDuckingPreference(asc2 *IAudioSessionControl2, optOut bool) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
