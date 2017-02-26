package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func main() {
	err := run()
	fmt.Println(err)
	return
}

func run() (err error) {
	ole.CoInitialize(0)
	unknown, err := ole.CreateInstance(CLSID_MMDeviceEnumerator, nil)
	if err != nil {
		return
	}
	defer unknown.Release()
	fmt.Println("@@@1")

	var de *IMMDeviceEnumerator
	if err = unknown.PutQueryInterface(IID_IMMDeviceEnumerator, &de); err != nil {
		return
	}
	defer de.Release()
	fmt.Println("@@@2")

	var dc *IMMDeviceCollection
	if err = de.EnumAudioEndpoints(ERender, DEVICE_STATE_ACTIVE, &dc); err != nil {
		return
	}
	fmt.Println("@@@3")

	var count uint32
	if err = dc.GetCount(&count); err != nil {
		return
	}
	fmt.Printf("%d devices found\n", count)

	var mmd *IMMDevice
	if err = dc.Item(count-1, &mmd); err != nil {
		return
	}
	defer mmd.Release()

	var strId uint32
	if err = mmd.GetId(&strId); err != nil {
		return
	}
	fmt.Println(strId)
	start := unsafe.Pointer(uintptr(strId))
	var str []uint16
	var i int
	for {
		item := *(*uint16)(unsafe.Pointer(uintptr(start) + 4*uintptr(i)))
		if item == 0 {
			break
		}
		str = append(str, item)
		fmt.Printf("%d ", item)
		i += 1
	}
	fmt.Printf("\n")
	fmt.Println(syscall.UTF16ToString(str))

	PKEY_Device_FriendlyName := DefinePropertyKey(0xa45c254e, 0xdf1c, 0x4efd, 0x80, 0x20, 0x67, 0xd1, 0x46, 0xa8, 0x50, 0xe0, 14)
	fmt.Println(PKEY_Device_FriendlyName)
	return
}
