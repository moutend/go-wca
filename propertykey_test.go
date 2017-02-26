package main

import (
	"testing"
)

func TestDefinePropertyKey(t *testing.T) {
	PKEY_Device_FriendlyName := DefinePropertyKey(0xa45c254e, 0xdf1c, 0x4efd, 0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0, 14) // DEVPROP_TYPE_STRING
	fmt.Println(PKEY_Device_FriendlyName)
	return
}
