package main

import (
	"fmt"

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
	err = unknown.PutQueryInterface(IID_IMMDeviceEnumerator, &de)
	if err != nil {
		return
	}
	defer de.Release()
	fmt.Println("@@@2")

	var dc *IMMDeviceCollection
	err = de.EnumAudioEndpoints(ERender, DEVICE_STATE_ACTIVE, &dc)
	if err != nil {
		return
	}
	fmt.Println("@@@3")

	var count uint32
	err = dc.GetCount(&count)
	if err != nil {
	}
	fmt.Printf("%d devices found\n", count)

	var mmd *IMMDevice
	err = dc.Item(count-1, &mmd)
	if err != nil {
		return
	}
	var state uint32
	err = mmd.GetState(&state)
	if err != nil {
		return
	}
	fmt.Printf("%d is %s\n", count-1, StringifyState(state))

	var aev *IAudioEndpointVolume
	if err = mmd.Activate(IID_IAudioEndpointVolume, CLSCTX_INPROC_SERVER, nil, &aev); err != nil {
		return
	}
	fmt.Println("@@@@")

	var channelCount uint32
	if err = aev.GetChannelCount(&channelCount); err != nil {
		return
	}
	fmt.Printf("device has %d channels\n", channelCount)
	return
}
