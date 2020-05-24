package wca

import (
	"github.com/go-ole/go-ole"
)

var (
	// Core Audio Interfaces - MSDN
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd370805(v=vs.85).aspx
	// The following IIDs are defined in mmdeviceapi.h.
	IID_IMMNotificationClient = ole.NewGUID("{7991EEC9-7E89-4D85-8390-6C703CEC60C0}")
	IID_IMMDevice             = ole.NewGUID("{D666063F-1587-4E43-81F1-B948E807363F}")
	IID_IMMDeviceCollection   = ole.NewGUID("{0BD7A1BE-7A1A-44DB-8397-CC5392387B5E}")
	IID_IMMEndpoint           = ole.NewGUID("{1BE09788-6894-4089-8586-9A2A6C265AC5}")
	IID_IMMDeviceEnumerator   = ole.NewGUID("{A95664D2-9614-4F35-A746-DE8DB63617E6}")
	IID_IMMDeviceActivator    = ole.NewGUID("{3B0D0EA4-D0A9-4B0E-935B-09516746FAC0}")
	CLSID_MMDeviceEnumerator  = ole.NewGUID("{BCDE0395-E52F-467C-8E3D-C4579291692E}")

	// The following IIDs are defined in Audioclient.h.
	IID_IAudioClient          = ole.NewGUID("{1CB9AD4C-DBFA-4c32-B178-C2F568A703B2}")
	IID_IAudioClient2         = ole.NewGUID("{726778CD-F60A-4eda-82DE-E47610CD78AA}")
	IID_IAudioClient3         = ole.NewGUID("{7ED4EE07-8E67-4CD4-8C1A-2B7A5987AD42}")
	IID_IAudioRenderClient    = ole.NewGUID("{F294ACFC-3146-4483-A7BF-ADDCA7C260E2}")
	IID_IAudioCaptureClient   = ole.NewGUID("{C8ADBD64-E71E-48a0-A4DE-185C395CD317}")
	IID_IAudioClock           = ole.NewGUID("{CD63314F-3FBA-4a1b-812C-EF96358728E7}")
	IID_IAudioClock2          = ole.NewGUID("{6f49ff73-6727-49ac-a008-d98cf5e70048}")
	IID_IAudioClockAdjustment = ole.NewGUID("{f6e4c0a0-46d9-4fb8-be21-57a3ef2b626c}")
	IID_ISimpleAudioVolume    = ole.NewGUID("{87CE5498-68D6-44E5-9215-6DA47EF883D8}")
	IID_IAudioStreamVolume    = ole.NewGUID("{93014887-242D-4068-8A15-CF5E93B90FE3}")
	IID_IChannelAudioVolume   = ole.NewGUID("{1C158861-B533-4B30-B1CF-E853E51C59B8}")

	// The following IDs are defined in audiopolicy.h.
	IID_IAudioSessionEvents          = ole.NewGUID("{24918ACC-64B3-37C1-8CA9-74A66E9957A8}")
	IID_IAudioSessionControl         = ole.NewGUID("{F4B1A599-7266-4319-A8CA-E70ACB11E8CD}")
	IID_IAudioSessionControl2        = ole.NewGUID("{bfb7ff88-7239-4fc9-8fa2-07c950be9c6d}")
	IID_IAudioSessionManager         = ole.NewGUID("{BFA971F1-4D5E-40BB-935E-967039BFBEE4}")
	IID_IAudioVolumeDuckNotification = ole.NewGUID("{C3B284D4-6D39-4359-B3CF-B56DDB3BB39C}")
	IID_IAudioSessionNotification    = ole.NewGUID("{641DD20B-4D41-49CC-ABA3-174B9477BB08}")
	IID_IAudioSessionEnumerator      = ole.NewGUID("{E2F5BB11-0570-40CA-ACDD-3AA01277DEE8}")
	IID_IAudioSessionManager2        = ole.NewGUID("{77AA99A0-1BD6-484F-8BC7-2C654C9A9B6F}")

	// The following IIDs are defined in devicetopology.h.
	IID_IKsControl              = ole.NewGUID("{28F54685-06FD-11D2-B27A-00A0C9223196}")
	IID_IPerChannelDbLevel      = ole.NewGUID("{C2F8E001-F205-4BC9-99BC-C13B1E048CCB}")
	IID_IAudioVolumeLevel       = ole.NewGUID("{7FB7B48F-531D-44A2-BCB3-5AD5A134B3DC}")
	IID_IAudioChannelConfig     = ole.NewGUID("{BB11C46F-EC28-493C-B88A-5DB88062CE98}")
	IID_IAudioLoudness          = ole.NewGUID("{7D8B1437-DD53-4350-9C1B-1EE2890BD938}")
	IID_IAudioInputSelector     = ole.NewGUID("{4F03DC02-5E6E-4653-8F72-A030C123D598}")
	IID_IAudioOutputSelector    = ole.NewGUID("{BB515F69-94A7-429e-8B9C-271B3F11A3AB}")
	IID_IAudioMute              = ole.NewGUID("{DF45AEEA-B74A-4B6B-AFAD-2366B6AA012E}")
	IID_IAudioBass              = ole.NewGUID("{A2B1A1D9-4DB3-425D-A2B2-BD335CB3E2E5}")
	IID_IAudioMidrange          = ole.NewGUID("{5E54B6D7-B44B-40D9-9A9E-E691D9CE6EDF}")
	IID_IAudioTreble            = ole.NewGUID("{0A717812-694E-4907-B74B-BAFA5CFDCA7B}")
	IID_IAudioAutoGainControl   = ole.NewGUID("{85401FD4-6DE4-4b9d-9869-2D6753A82F3C}")
	IID_IAudioPeakMeter         = ole.NewGUID("{DD79923C-0599-45e0-B8B6-C8DF7DB6E796}")
	IID_IDeviceSpecificProperty = ole.NewGUID("{3B22BCBF-2586-4af0-8583-205D391B807C}")
	IID_IKsFormatSupport        = ole.NewGUID("{3CB4A69D-BB6F-4D2B-95B7-452D2C155DB5}")
	IID_IKsJackDescription      = ole.NewGUID("{4509F757-2D46-4637-8E62-CE7DB944F57B}")
	IID_IKsJackDescription2     = ole.NewGUID("{478F3A9B-E0C9-4827-9228-6F5505FFE76A}")
	IID_IKsJackSinkInformation  = ole.NewGUID("{D9BD72ED-290F-4581-9FF3-61027A8FE532}")
	IID_IPartsList              = ole.NewGUID("{6DAA848C-5EB0-45CC-AEA5-998A2CDA1FFB}")
	IID_IPart                   = ole.NewGUID("{AE2DE0E4-5BCA-4F2D-AA46-5D13F8FDB3A9}")
	IID_IConnector              = ole.NewGUID("{9c2c4058-23f5-41de-877a-df3af236a09e}")
	IID_ISubunit                = ole.NewGUID("{82149A85-DBA6-4487-86BB-EA8F7FEFCC71}")
	IID_IControlInterface       = ole.NewGUID("{45d37c3f-5140-444a-ae24-400789f3cbf3}")
	IID_IControlChangeNotify    = ole.NewGUID("{A09513ED-C709-4d21-BD7B-5F34C47F3947}")
	IID_IDeviceTopology         = ole.NewGUID("{2A07407E-6497-4A18-9787-32F79BD0D98F}")
	CLSID_DeviceTopology        = ole.NewGUID("{1DF639D0-5EC1-47AA-9379-828DC1AA8C59}")

	// The following IIDs are defined in endpointvolume.h.
	IID_IAudioEndpointVolumeCallback = ole.NewGUID("{657804FA-D6AD-4496-8A60-352752AF4F89}")
	IID_IAudioEndpointVolume         = ole.NewGUID("{5CDF2C82-841E-4546-9722-0CF74078229A}")
	IID_IAudioEndpointVolumeEx       = ole.NewGUID("{66E11784-F695-4F28-A505-A7080081A78F}")
	IID_IAudioMeterInformation       = ole.NewGUID("{C02216F6-8C67-4B5B-9D00-D008E73E0064}")
)
