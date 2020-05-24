// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func accGetBuffer(acc *IAudioCaptureClient, data **byte, framesToRead, flags *uint32, devicePosition, qpcPosition *uint64) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func accReleaseBuffer(acc *IAudioCaptureClient, framesRead uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func accGetNextPacketSize(acc *IAudioCaptureClient, framesInNextPacket *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
