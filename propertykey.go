package main

type PropertyKey struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
	PID   uint32
}

func DefinePropertyKey(l uint32, w1, w2 uint16, b1, b2, b3, b4, b5, b6, b7, b8 byte, pid uint32) *PropertyKey {
	p := PropertyKey{
		Data1: l,
		Data2: w1,
		Data3: w2,
		Data4: [8]byte{b1, b2, b3, b4, b5, b6, b7, b8},
		PID:   pid,
	}

	return &p
}
