package main

import (
	"github.com/go-ole/go-ole"
)

var (
	IID_IMMDeviceEnumerator          = ole.NewGUID("{A95664D2-9614-4F35-A746-DE8DB63617E6}")
	CLSID_MMDeviceEnumerator         = ole.NewGUID("{BCDE0395-E52F-467C-8E3D-C4579291692E}")
	IID_IAudioEndpointVolumeCallback = ole.NewGUID("{657804FA-D6AD-4496-8A60-352752AF4F89}")
	IID_IAudioEndpointVolume         = ole.NewGUID("{5CDF2C82-841E-4546-9722-0CF74078229A}")
	IID_IAudioEndpointVolumeEx       = ole.NewGUID("{66E11784-F695-4F28-A505-A7080081A78F}")
	IID_IAudioMeterInformation       = ole.NewGUID("{C02216F6-8C67-4B5B-9D00-D008E73E0064}")
)
