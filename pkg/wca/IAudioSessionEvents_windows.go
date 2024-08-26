//go:build windows
// +build windows

package wca

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func aseQueryInterface(this uintptr, riid *ole.GUID, ppInterface *uintptr) int64 {
	*ppInterface = 0

	if ole.IsEqualGUID(riid, ole.IID_IUnknown) ||
		ole.IsEqualGUID(riid, IID_IAudioSessionEvents) {
		aseAddRef(this)
		*ppInterface = this

		return ole.S_OK
	}

	return ole.E_NOINTERFACE
}

func aseAddRef(this uintptr) int64 {
	ase := (*IAudioSessionNotification)(unsafe.Pointer(this))

	ase.refCount += 1

	return int64(ase.refCount)
}

func aseRelease(this uintptr) int64 {
	ase := (*IAudioSessionNotification)(unsafe.Pointer(this))

	ase.refCount -= 1

	return int64(ase.refCount)
}

func aseOnDisplayNameChanged(this, lpcwstrNewDisplayName, lpcguidEventContext uintptr) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnDisplayNameChanged == nil {
		return ole.S_OK
	}

	newDisplayName := LPCWSTRToString(lpcwstrNewDisplayName, 1024)
	eventContext := (*ole.GUID)(unsafe.Pointer(lpcguidEventContext))

	err := ase.callback.OnDisplayNameChanged(newDisplayName, eventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnIconPathChanged(this, lpcwstrNewIconPath, lpcguidEventContext uintptr) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnIconPathChanged == nil {
		return ole.S_OK
	}

	newIconPath := LPCWSTRToString(lpcwstrNewIconPath, 1024)
	eventContext := (*ole.GUID)(unsafe.Pointer(lpcguidEventContext))

	err := ase.callback.OnIconPathChanged(newIconPath, eventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnSimpleVolumeChanged(this uintptr, newVolume float32, muteValue uint32, lpcguidEventContext uintptr) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnSimpleVolumeChanged == nil {
		return ole.S_OK
	}

	mute := false
	if muteValue > 0 {
		mute = true
	}
	eventContext := (*ole.GUID)(unsafe.Pointer(lpcguidEventContext))

	err := ase.callback.OnSimpleVolumeChanged(newVolume, mute, eventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnChannelVolumeChanged(this uintptr, channelCount uint32, newChannelVolumeArray []float32, changedChannel uint32, lpcguidEventContext uintptr) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnChannelVolumeChanged == nil {
		return ole.S_OK
	}

	eventContext := (*ole.GUID)(unsafe.Pointer(lpcguidEventContext))

	err := ase.callback.OnChannelVolumeChanged(int(changedChannel), newChannelVolumeArray, int(changedChannel), eventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnGroupingParamChanged(this, lpcguidNewGroupingParam, lpcguidEventContext uintptr) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnGroupingParamChanged == nil {
		return ole.S_OK
	}

	newGroupingParam := (*ole.GUID)(unsafe.Pointer(lpcguidNewGroupingParam))
	eventContext := (*ole.GUID)(unsafe.Pointer(lpcguidEventContext))

	err := ase.callback.OnGroupingParamChanged(newGroupingParam, eventContext)

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnStateChanged(this uintptr, newState int64) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnStateChanged == nil {
		return ole.S_OK
	}

	err := ase.callback.OnStateChanged(AudioSessionState(newState))

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func aseOnSessionDisconnected(this uintptr, disconnectReason int64) int64 {
	ase := (*IAudioSessionEvents)(unsafe.Pointer(this))

	if ase.callback.OnSessionDisconnected == nil {
		return ole.S_OK
	}

	err := ase.callback.OnSessionDisconnected(AudioSessionDisconnectReason(disconnectReason))

	if err != nil {
		return ole.E_FAIL
	}

	return ole.S_OK
}

func NewIAudioSessionEvents(callback IAudioSessionEventsCallback) *IAudioSessionEvents {
	vTable := &IAudioSessionEventsVtbl{}

	if callback.OnSimpleVolumeChanged != nil || callback.OnChannelVolumeChanged != nil {
		panic("float arguments not supported, see https://github.com/golang/go/issues/45300")
	}

	// IUnknown methods
	vTable.QueryInterface = syscall.NewCallback(aseQueryInterface)
	vTable.AddRef = syscall.NewCallback(aseAddRef)
	vTable.Release = syscall.NewCallback(aseRelease)

	// IAudioSessionEvents methods
	vTable.OnDisplayNameChanged = syscall.NewCallback(aseOnDisplayNameChanged)
	vTable.OnIconPathChanged = syscall.NewCallback(aseOnIconPathChanged)
	// vTable.OnSimpleVolumeChanged = syscall.NewCallback(aseOnSimpleVolumeChanged)
	// vTable.OnChannelVolumeChanged = syscall.NewCallback(aseOnChannelVolumeChanged)
	vTable.OnGroupingParamChanged = syscall.NewCallback(aseOnGroupingParamChanged)
	vTable.OnStateChanged = syscall.NewCallback(aseOnStateChanged)
	vTable.OnSessionDisconnected = syscall.NewCallback(aseOnSessionDisconnected)

	ase := &IAudioSessionEvents{}

	ase.vTable = vTable
	ase.callback = callback

	return ase
}
