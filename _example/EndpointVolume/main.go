package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

var version = "latest"
var revision = "latest"

type GainFlag struct {
	Value float32
	IsSet bool
}

func (f *GainFlag) Set(value string) (err error) {
	var v float64
	if v, err = strconv.ParseFloat(value, 64); err != nil {
		return
	}
	if v > 0 {
		err = fmt.Errorf("gain must be less than or equal to 0 [dB]")
	}
	f.Value = float32(v)
	f.IsSet = true
	return
}
func (f GainFlag) String() string {
	return fmt.Sprintf("%v", f.Value)
}

type VolumeFlag struct {
	Value float32
	IsSet bool
}

func (f *VolumeFlag) Set(value string) (err error) {
	var v float64
	if v, err = strconv.ParseFloat(value, 64); err != nil {
		return
	}
	if v > 1.0 || v < 0.0 {
		err = fmt.Errorf("volume range is 0.0 to 1.0")
		return
	}
	f.Value = float32(v)
	f.IsSet = true
	return
}

func (f *VolumeFlag) String() string {
	return fmt.Sprintf("%v", f.Value)
}

type MuteFlag struct {
	Value bool
	IsSet bool
}

func (f *MuteFlag) Set(value string) (err error) {
	if value != "true" && value != "false" {
		err = fmt.Errorf("set 'true' or 'false'")
		return
	}
	if value == "true" {
		f.Value = true
	}
	f.IsSet = true
	return
}

func (f *MuteFlag) String() string {
	return fmt.Sprintf("%v", f.Value)
}

func main() {
	var err error
	if err = run(os.Args); err != nil {
		log.Fatal(err)
	}
	return
}

func run(args []string) (err error) {
	var volumeFlag VolumeFlag
	var gainFlag GainFlag
	var muteFlag MuteFlag
	var versionFlag bool

	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	f.Var(&volumeFlag, "volume", "Specify volume as a scalar value")
	f.Var(&volumeFlag, "v", "Alias of ---volume")
	f.Var(&gainFlag, "gain", "Specify volume as a gain value [dB]")
	f.Var(&gainFlag, "g", "Alias of --gain")
	f.Var(&muteFlag, "mute", "Specify mute state (default is false)")
	f.Var(&muteFlag, "m", "Alias of --mute")
	f.BoolVar(&versionFlag, "version", false, "Show version")
	f.Parse(args[1:])

	if versionFlag {
		fmt.Printf("%s-%s\n", version, revision)
		return
	}
	if err = endpointVolume(gainFlag, volumeFlag, muteFlag); err != nil {
		return
	}
	fmt.Println("Successfully done")
	return
}

func endpointVolume(gainFlag GainFlag, volumeFlag VolumeFlag, muteFlag MuteFlag) (err error) {
	fmt.Println(volumeFlag)
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}
	defer ole.CoUninitialize()

	var mmde *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err = mmde.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
		return
	}
	defer mmd.Release()

	var ps *wca.IPropertyStore
	if err = mmd.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		return
	}
	defer ps.Release()

	var pv wca.PROPVARIANT
	if err = ps.GetValue(&wca.PKEY_Device_FriendlyName, &pv); err != nil {
		return
	}
	fmt.Printf("%s\n", pv.String())

	var aev *wca.IAudioEndpointVolume
	if err = mmd.Activate(wca.IID_IAudioEndpointVolume, wca.CLSCTX_ALL, nil, &aev); err != nil {
		return
	}
	defer aev.Release()

	if gainFlag.IsSet {
		fmt.Println("setting gain flag")
		if err = aev.SetMasterVolumeLevel(gainFlag.Value, nil); err != nil {
			return
		}
	}
	if volumeFlag.IsSet {
		if err = aev.SetMasterVolumeLevelScalar(volumeFlag.Value, nil); err != nil {
			return
		}
	}
	if muteFlag.IsSet {
		if err = aev.SetMute(muteFlag.Value, nil); err != nil {
			return
		}
	}
	var channels uint32
	if err = aev.GetChannelCount(&channels); err != nil {
		return
	}

	var mute bool
	if err = aev.GetMute(&mute); err != nil {
		return
	}

	var masterVolumeLevel float32
	if err = aev.GetMasterVolumeLevel(&masterVolumeLevel); err != nil {
		return
	}

	var masterVolumeLevelScalar float32
	if err = aev.GetMasterVolumeLevelScalar(&masterVolumeLevelScalar); err != nil {
		return
	}

	fmt.Println("--------")
	fmt.Printf("Channels: %d\n", channels)
	fmt.Printf("Mute state: %v\n", mute)
	fmt.Println("Master volume level:")
	fmt.Printf("  %v [dB]\n", masterVolumeLevel)
	fmt.Printf("  %v [scalar]\n", masterVolumeLevelScalar)
	fmt.Println("--------")

	return
}
