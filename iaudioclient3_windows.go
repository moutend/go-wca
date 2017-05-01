// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func ac3GetSharedModeEnginePeriod(ac3 *IAudioClient3, wfx *WAVEFORMATEX, defaultPeriodInFrames, fundamentalPeriodInFrames, minPeriodInFrames, maxPeriodInFrames *uint32) (err error) {
	hr, _, _ := syscall.Syscall6(
		ac3.VTable().GetSharedModeEnginePeriod,
		6,
		uintptr(unsafe.Pointer(ac3)),
		uintptr(unsafe.Pointer(wfx)),
		uintptr(unsafe.Pointer(defaultPeriodInFrames)),
		uintptr(unsafe.Pointer(fundamentalPeriodInFrames)),
		uintptr(unsafe.Pointer(minPeriodInFrames)),
		uintptr(unsafe.Pointer(maxPeriodInFrames)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ac3GetCurrentSharedModeEnginePeriod(ac3 *IAudioClient3, wfx **WAVEFORMATEX, currentPeriodInFrames *uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		ac3.VTable().GetCurrentSharedModeEnginePeriod,
		3,
		uintptr(unsafe.Pointer(ac3)),
		uintptr(unsafe.Pointer(wfx)),
		uintptr(unsafe.Pointer(currentPeriodInFrames)))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func ac3InitializeSharedAudioStream(ac3 *IAudioClient3, streamFlags, periodInFrames uint32, wfx *WAVEFORMATEX, audioSessionGUID *ole.GUID) (err error) {
	hr, _, _ := syscall.Syscall6(
		ac3.VTable().InitializeSharedAudioStream,
		5,
		uintptr(unsafe.Pointer(ac3)),
		uintptr(streamFlags),
		uintptr(periodInFrames),
		uintptr(unsafe.Pointer(wfx)),
		uintptr(unsafe.Pointer(audioSessionGUID)),
		0)
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
