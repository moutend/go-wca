package wca

// ERole represents ERole enumeration.
//
// For more details, visit the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/api/mmdeviceapi/ne-mmdeviceapi-erole
type ERole uint64

// EDataFlow represents EDataFlow enumeration.
//
// For more details, visit the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/api/mmdeviceapi/ne-mmdeviceapi-edataflow
type EDataFlow uint64

// REFERENCE_TIME represents REFERENCE_TIME data type.
//
// For more details, visit the MSDN.
//
// https://docs.microsoft.com/en-us/windows/win32/directshow/reference-time
type REFERENCE_TIME int64

// AudioSessionDisconnectReason represents AudioSessionDisconnectReason enumeration.
//
// For more details, visit the MSDN.
//
// https://learn.microsoft.com/en-us/windows/win32/api/audiopolicy/nf-audiopolicy-iaudiosessionevents-onsessiondisconnected
type AudioSessionDisconnectReason int64

// AudioSessionState represents AudioSessionState enumeration.
//
// For more details, visit the MSDN.
//
// https://learn.microsoft.com/en-us/windows/win32/api/audiopolicy/nf-audiopolicy-iaudiosessionevents-onstatechanged
// https://learn.microsoft.com/en-us/windows/win32/api/audiosessiontypes/ne-audiosessiontypes-audiosessionstate
type AudioSessionState int64
