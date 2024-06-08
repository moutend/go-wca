package wca

import "github.com/go-ole/go-ole"

// WAVEFORMATEX structure defines the format of waveform-audio data
// https://learn.microsoft.com/en-us/windows/win32/api/mmreg/ns-mmreg-waveformatex
type WAVEFORMATEX struct {
	WFormatTag      uint16
	NChannels       uint16
	NSamplesPerSec  uint32
	NAvgBytesPerSec uint32
	NBlockAlign     uint16
	WBitsPerSample  uint16
	CbSize          uint16
}

// WAVEFORMATEXTENSIBLE structure extends WAVEFORMATEX
// See: https://learn.microsoft.com/en-us/windows/win32/api/mmreg/ns-mmreg-waveformatextensible
type WAVEFORMATEXTENSIBLE struct {
	Format        WAVEFORMATEX
	Samples       uint16 // union { ValidBitsPerSample or SamplesPerBlock or Reserved }
	DwChannelMask uint32
	SubFormat     ole.GUID
}

var (
	// KSDATAFORMAT_SUBTYPE_PCM PCM Audio Data Format
	KSDATAFORMAT_SUBTYPE_PCM = &ole.GUID{
		Data1: 0x00000001,
		Data2: 0x0000,
		Data3: 0x0010,
		Data4: [8]byte{0x80, 0x00, 0x00, 0xaa, 0x00, 0x38, 0x9b, 0x71},
	}

	// KSDATAFORMAT_SUBTYPE_IEEE_FLOAT IEEE Float Audio Data Format
	KSDATAFORMAT_SUBTYPE_IEEE_FLOAT = &ole.GUID{
		Data1: 0x00000003,
		Data2: 0x0000,
		Data3: 0x0010,
		Data4: [8]byte{0x80, 0x00, 0x00, 0xaa, 0x00, 0x38, 0x9b, 0x71},
	}

	// KSDATAFORMAT_SUBTYPE_ALAW ALAW Audio Data Format
	KSDATAFORMAT_SUBTYPE_ALAW = &ole.GUID{
		Data1: 0x00000006,
		Data2: 0x0000,
		Data3: 0x0010,
		Data4: [8]byte{0x80, 0x00, 0x00, 0xaa, 0x00, 0x38, 0x9b, 0x71},
	}

	// KSDATAFORMAT_SUBTYPE_MULAW MULAW Audio Data Format
	KSDATAFORMAT_SUBTYPE_MULAW = &ole.GUID{
		Data1: 0x00000007,
		Data2: 0x0000,
		Data3: 0x0010,
		Data4: [8]byte{0x80, 0x00, 0x00, 0xaa, 0x00, 0x38, 0x9b, 0x71},
	}
)
