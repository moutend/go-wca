package main

import (
	"fmt"

	"github.com/go-ole/go-ole"
)

func main() {
	ole.CoInitialize(0)
	unknown, err := ole.CreateInstance(CLSID_MMDeviceEnumerator, nil)
	if err != nil {
		return
	}
	defer unknown.Release()
	fmt.Println("@@@1")
	var de *IMMDeviceEnumerator
	err = unknown.PutQueryInterface(IID_IMMDeviceEnumerator, &de)
	if err != nil {
		return
	}
	defer de.Release()
	fmt.Println("@@@2")
	var dc *IMMDeviceCollection
	err = de.EnumAudioEndpoints(ERender, DEVICE_STATE_ACTIVE, &dc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("@@@3")
	var count uint32
	err = dc.GetCount(&count)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}
