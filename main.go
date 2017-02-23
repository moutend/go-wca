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
	defer aev.Release()
	fmt.Println("@@@@")

	var channelCount uint32
	if err = aev.GetChannelCount(&channelCount); err != nil {
		return
	}
	fmt.Printf("device has %d channels\n", channelCount)

	var level float32
	level = -1.234
	if err = aev.GetMasterVolumeLevelScalar(&level); err != nil {
		return
	}
	fmt.Printf("current volume is %f\n", level)
	level = 0.5
	if err = aev.VolumeStepUp(nil); err != nil {
		return
	}
	fmt.Println("Volume step up")

	if err = aev.VolumeStepDown(nil); err != nil {
		return
	}
	fmt.Println("Volume step down")

	//if err = aev.SetMasterVolumeLevelScalar(level, nil); err != nil {
	//	return
	//}
	var minDB float32
	var maxDB float32
	var incrementDB float32
	if err = aev.GetVolumeRange(&minDB, &maxDB, &incrementDB); err != nil {
		return
	}
	fmt.Printf("%f to %f (step: %f)\n", minDB, maxDB, incrementDB)
	var step uint32
	var stepCount uint32
	if err = aev.GetVolumeStepInfo(&step, &stepCount); err != nil {
		return
	}
	fmt.Println(step, stepCount)
	fmt.Println("done")
	return
}
