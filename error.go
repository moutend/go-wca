package wca

import (
	"fmt"

	"github.com/go-ole/go-ole"
)

const (
	AUDCLNT_E_NOT_INITIALIZED              = 0x001
	AUDCLNT_E_ALREADY_INITIALIZED          = 0x002
	AUDCLNT_E_WRONG_ENDPOINT_TYPE          = 0x003
	AUDCLNT_E_DEVICE_INVALIDATED           = 0x004
	AUDCLNT_E_NOT_STOPPED                  = 0x005
	AUDCLNT_E_BUFFER_TOO_LARGE             = 0x006
	AUDCLNT_E_OUT_OF_ORDER                 = 0x007
	AUDCLNT_E_UNSUPPORTED_FORMAT           = 0x008
	AUDCLNT_E_INVALID_SIZE                 = 0x009
	AUDCLNT_E_DEVICE_IN_USE                = 0x00a
	AUDCLNT_E_BUFFER_OPERATION_PENDING     = 0x00b
	AUDCLNT_E_THREAD_NOT_REGISTERED        = 0x00c
	AUDCLNT_E_EXCLUSIVE_MODE_NOT_ALLOWED   = 0x00e
	AUDCLNT_E_ENDPOINT_CREATE_FAILED       = 0x00f
	AUDCLNT_E_SERVICE_NOT_RUNNING          = 0x010
	AUDCLNT_E_EVENTHANDLE_NOT_EXPECTED     = 0x011
	AUDCLNT_E_EXCLUSIVE_MODE_ONLY          = 0x012
	AUDCLNT_E_BUFDURATION_PERIOD_NOT_EQUAL = 0x013
	AUDCLNT_E_EVENTHANDLE_NOT_SET          = 0x014
	AUDCLNT_E_INCORRECT_BUFFER_SIZE        = 0x015
	AUDCLNT_E_BUFFER_SIZE_ERROR            = 0x016
	AUDCLNT_E_CPUUSAGE_EXCEEDED            = 0x017
	AUDCLNT_E_BUFFER_ERROR                 = 0x018
	AUDCLNT_E_BUFFER_SIZE_NOT_ALIGNED      = 0x019
	AUDCLNT_E_INVALID_DEVICE_PERIOD        = 0x020
)

func NewError(hr uintptr) (err error) {
	fmt.Printf("hr is %d\n", hr)
	switch hr {
	case AUDCLNT_E_NOT_INITIALIZED:
		err = fmt.Errorf("audio stream has not been successfully initialized")
	case AUDCLNT_E_ALREADY_INITIALIZED:
		err = fmt.Errorf("IAudioClient object is already initialized")
	case AUDCLNT_E_WRONG_ENDPOINT_TYPE:
		err = fmt.Errorf("specifying an IAudioCaptureClient interface on a rendering endpoint or an IAudioRenderClient interface on a capture endpoint is not valid")
	case AUDCLNT_E_DEVICE_INVALIDATED:
		err = fmt.Errorf("audio endpoint device is unplugged or unavailable")
	case AUDCLNT_E_NOT_STOPPED:
		err = fmt.Errorf("audio stream must be stopped before starting")
	case AUDCLNT_E_BUFFER_TOO_LARGE:
		err = fmt.Errorf("exceeds available buffer space")
	case AUDCLNT_E_OUT_OF_ORDER:
		err = fmt.Errorf("previous IAudioRenderClient::GetBuffer call is still in effect")
	case AUDCLNT_E_UNSUPPORTED_FORMAT:
		err = fmt.Errorf("audio engine (shared mode) or audio endpoint device (exclusive mode) does not support the specified format")
	case AUDCLNT_E_INVALID_SIZE:
		err = fmt.Errorf("written frame size exceeds requested frame size ")
	case AUDCLNT_E_DEVICE_IN_USE:
		err = fmt.Errorf("endpoint device is already in use")
	case AUDCLNT_E_BUFFER_OPERATION_PENDING:
		err = fmt.Errorf("buffer cannot be accessed while reset is in progress")
	case AUDCLNT_E_THREAD_NOT_REGISTERED:
		err = fmt.Errorf("")
	case AUDCLNT_E_EXCLUSIVE_MODE_NOT_ALLOWED:
		err = fmt.Errorf("user has disabled exclusive-mode use of the device")
	case AUDCLNT_E_ENDPOINT_CREATE_FAILED:
		err = fmt.Errorf("failed to create the audio endpoint for the render or the capture device")
	case AUDCLNT_E_SERVICE_NOT_RUNNING:
		err = fmt.Errorf("Windows audio service is not running")
	case AUDCLNT_E_EVENTHANDLE_NOT_EXPECTED:
		fmt.Errorf("")
	case AUDCLNT_E_EXCLUSIVE_MODE_ONLY:
		fmt.Errorf("")
	case AUDCLNT_E_BUFDURATION_PERIOD_NOT_EQUAL:
		fmt.Errorf("duration and peridoicity is not equal")
	case AUDCLNT_E_EVENTHANDLE_NOT_SET:
		err = fmt.Errorf("audio stream is configured to use event driven buffering but event is not set")
	case AUDCLNT_E_INCORRECT_BUFFER_SIZE:
		err = fmt.Errorf("")
	case AUDCLNT_E_BUFFER_SIZE_ERROR:
		err = fmt.Errorf("requested buffer size is out of range")
	case AUDCLNT_E_CPUUSAGE_EXCEEDED:
		err = fmt.Errorf("exceeded maximum CPU usage")
	case AUDCLNT_E_BUFFER_ERROR:
		err = fmt.Errorf("")
	case AUDCLNT_E_BUFFER_SIZE_NOT_ALIGNED:
		err = fmt.Errorf("requested buffer size is not aligned")
	case AUDCLNT_E_INVALID_DEVICE_PERIOD:
		err = fmt.Errorf("device period requested by an exclusive-mode client is greater than 500 milliseconds")
	default:
		err = ole.NewError(hr)
	}
	return
}
