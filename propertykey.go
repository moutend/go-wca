package main

import (
	"github.com/go-ole/go-ole"
)

type PropertyKey struct {
	ole.GUID
	PID uint32
}

func NewPropertyKey(l uint32, w1, w2 uint16, b1, b2, b3, b4, b5, b6, b7, b8 byte, pid uint32) PropertyKey {
	g := ole.GUID{
		Data1: l,
		Data2: w1,
		Data3: w2,
		Data4: [8]byte{b1, b2, b3, b4, b5, b6, b7, b8},
	}
	return PropertyKey{g, pid}
}
