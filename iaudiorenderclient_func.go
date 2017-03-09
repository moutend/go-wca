// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func arcGetBuffer(arc *IAudioRenderClient, requiredBufferSize uint32, data interface{}) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func arcReleaseBuffer(arc *IAudioRenderClient, writtenBufferSize, flag uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
