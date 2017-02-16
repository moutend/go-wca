package main

const (
	DEVICE_STATE_ACTIVE     = 0x00000001
	DEVICE_STATE_DISABLED   = 0x00000002
	DEVICE_STATE_NOTPRESENT = 0x00000004
	DEVICE_STATE_UNPLUGGED  = 0x00000008
	DEVICE_STATEMASK_ALL    = 0x0000000F
	ERender                 = 0x1
	ECapture                = 0x2
	EAll                    = 0x3
	EDataFlow_enum_count    = 0x4
)
