// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func acStart(ac *IAudioClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acStop(ac *IAudioClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func acReset(ac *IAudioClient) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
