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
	var me *IMMDeviceEnumerator
	err = unknown.PutQueryInterface(IID_IMMDeviceEnumerator, &me)
	if err != nil {
		return
	}
	defer me.Release()
	fmt.Println("@@@2")
	var dc *IMMDeviceCollection
	err = me.EnumAudioEndpoints(ERender, DEVICE_STATE_ACTIVE, &dc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("@@@3")
}
