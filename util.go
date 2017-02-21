package main

func StringifyState(state uint32) string {
	switch state {
	case DEVICE_STATE_ACTIVE:
		return "active"
	case DEVICE_STATE_DISABLED:
		return "disabled"
	default:
		return ""
	}
}
