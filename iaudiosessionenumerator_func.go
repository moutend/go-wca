// +build !windows

package wca

import "github.com/go-ole/go-ole"

func aseGetCount(ase *IAudioSessionEnumerator, sessionCount *int) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func aseGetSession(ase *IAudioSessionEnumerator, sessionCount int, session **IAudioSessionControl) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
