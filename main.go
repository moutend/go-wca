package main

import (
	"fmt"
	//"syscall"
	//"unsafe"

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

	var ps *IPropertyStore
	if err = mmd.OpenPropertyStore(STGM_READ, ps); err != nil {
		return
	}
	defer ps.Release()

	return
}
