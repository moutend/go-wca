// +build !windows

package wca

import "github.com/go-ole/go-ole"

func asmGetAudioSessionControl(asm *IAudioSessionManager, audioSessionGUID *ole.GUID, streamFlags uint32, sessionControl **IAudioSessionControl) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}

func asmGetSimpleAudioVolume(asm *IAudioSessionManager, audioSessionGUID *ole.GUID, streamFlags uint32, audioVolume **ISimpleAudioVolume) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
