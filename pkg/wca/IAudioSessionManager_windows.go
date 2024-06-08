//go:build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func asmGetAudioSessionControl(asm *IAudioSessionManager, audioSessionGUID *ole.GUID, streamFlags uint32, sessionControl **IAudioSessionControl) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm.VTable().GetAudioSessionControl,
		4,
		uintptr(unsafe.Pointer(asm)),
		uintptr(unsafe.Pointer(audioSessionGUID)),
		uintptr(streamFlags),
		uintptr(unsafe.Pointer(sessionControl)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func asmGetSimpleAudioVolume(asm *IAudioSessionManager, audioSessionGUID *ole.GUID, streamFlags uint32, audioVolume **ISimpleAudioVolume) (err error) {
	hr, _, _ := syscall.SyscallN(
		asm.VTable().GetSimpleAudioVolume,
		4,
		uintptr(unsafe.Pointer(asm)),
		uintptr(unsafe.Pointer(audioSessionGUID)),
		uintptr(streamFlags),
		uintptr(unsafe.Pointer(audioVolume)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
