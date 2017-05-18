package wca

type IMMNotificationClient struct {
	VTable *IMMNotificationClientVtbl
}

type IMMNotificationClientVtbl struct {
	QueryInterface         uintptr
	AddRef                 uintptr
	Release                uintptr
	OnDeviceStateChanged   uintptr
	OnDeviceAdded          uintptr
	OnDeviceRemoved        uintptr
	OnDefaultDeviceChanged uintptr
	OnPropertyValueChanged uintptr
}
