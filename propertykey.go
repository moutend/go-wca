package wca

import (
	"github.com/go-ole/go-ole"
)

type PropertyKey struct {
	ole.GUID
	PID uint32
}

var (
	PKEY_Device_FriendlyName = DefinePropertyKey(0xa45c254e, 0xdf1c, 0x4efd, 0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0, 14) // DEVPROP_TYPE_STRING
)

func DefinePropertyKey(l uint32, w1, w2 uint16, b1, b2, b3, b4, b5, b6, b7, b8 byte, pid uint32) PropertyKey {
	g := ole.GUID{
		Data1: l,
		Data2: w1,
		Data3: w2,
		Data4: [8]byte{b1, b2, b3, b4, b5, b6, b7, b8},
	}
	return PropertyKey{g, pid}
}
