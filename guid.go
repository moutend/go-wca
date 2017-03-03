package wca

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
	IID_IAudioClient                 = ole.NewGUID("{1CB9AD4C-DBFA-4c32-B178-C2F568A703B2}")
	IID_IAudioRenderClient           = ole.NewGUID("{F294ACFC-3146-4483-A7BF-ADDCA7C260E2}")
	IID_IAudioCaptureClient          = ole.NewGUID("{C8ADBD64-E71E-48a0-A4DE-185C395CD317}")
	IID_IAudioClock                  = ole.NewGUID("{CD63314F-3FBA-4a1b-812C-EF96358728E7}")
	IID_IAudioClock2                 = ole.NewGUID("{6f49ff73-6727-49ac-a008-d98cf5e70048}")
	IID_IAudioClockAdjustment        = ole.NewGUID("{f6e4c0a0-46d9-4fb8-be21-57a3ef2b626c}")
	IID_ISimpleAudioVolume           = ole.NewGUID("{87CE5498-68D6-44E5-9215-6DA47EF883D8}")
	IID_IAudioStreamVolume           = ole.NewGUID("{93014887-242D-4068-8A15-CF5E93B90FE3}")
	IID_IChannelAudioVolume          = ole.NewGUID("{1C158861-B533-4B30-B1CF-E853E51C59B8}")
)
