# go-wca

[![CircleCI](https://circleci.com/gh/moutend/go-wca/tree/master.svg?style=svg&circle-token=a9f340e1c5288df38bda7d816d38d4189893eafe)][status]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GitHub release](http://img.shields.io/github/release/moutend/go-wca.svg?style=flat-square)][release]

[status]: https://circleci.com/gh/moutend/go-wca/tree/master
[license]: https://github.com/moutend/go-wca/blob/master/LICENSE
[release]: https://github.com/moutend/go-wca/releases

Go bindings for Windows Core Audio API without using cgo.

This package allows you to do:

- Rendering audio with shared timer driven mode.
- Rendering audio with event driven mode.
- Capturing audio with shared timer driven mode.
- Capturing audio with event driven mode.
- Loopback capturing with shared timer mode.
- Loopback capturing with shared event mode.
- Rendering / Capturing with minimum buffer size (Windows 10 only)
- Change volume of master or each channels.
- Change volume of individual applications.

If you're not familiar with Windows Core Audio API, [the official documentation about Core Audio API on MSDN](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370802(v=vs.85).aspx) helps you to get started.

## Prerequisites

- Go 1.8 or later
- `go-ole` ([github.com/go-ole/go-ole](https://github.com/go-ole/go-ole))

## Examples

The examples are located in `example` directory. You can download [executables](https://github.com/moutend/go-wca/releases) or build by yourself. For more information, please read the README.md in each examples.

If you want to build all examples at once, run the command below:

```console
mkdir bin
make VERSION=latest
```

Then the executables are generated in `bin` directory.

## Documentation

Each APIs in this package correspond to native COM APIs, so that you can refer the documentation on MSDN as a full documentation of this package.

The following list contains the links to the documentation of native API which is available in this package.
- MMDevice API
  - [IMMDevice](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371395(v=vs.85).aspx)
  - [IMMDeviceCollection](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371396(v=vs.85).aspx "IMMDeviceCollection")
  - [IMMDeviceEnumerator](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371399(v=vs.85).aspx "IMMDeviceEnumerator")
  - [IMMEndpoint](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371414(v=vs.85).aspx "IMMEndpoint")
- Windows Audio Session API
  - [IAudioClient](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370865(v=vs.85).aspx "IAudioClient")
  - [IAudioClient2](https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx "IAudioClient2")
  - [IAudioClient3](https://msdn.microsoft.com/en-us/library/windows/desktop/dn911487(v=vs.85).aspx "IAudioClient3")
  - [IAudioCaptureClient](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370858(v=vs.85).aspx "IAudioCaptureClient")
  - [IAudioEndpointVolume](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370892(v=vs.85).aspx "IAudioEndpointVolume")
  - [IAudioRenderClient](https://msdn.microsoft.com/en-us/library/windows/desktop/dd368242(v=vs.85).aspx "IAudioRenderClient")
  - [IAudioSessionManager](https://docs.microsoft.com/en-us/windows/desktop/api/audiopolicy/nn-audiopolicy-iaudiosessionmanager "IAudioSessionManager")
  - [IAudioSessionManager2](https://docs.microsoft.com/en-us/windows/desktop/api/audiopolicy/nn-audiopolicy-iaudiosessionmanager2 "IAudioSessionManager2")
  - [IAudioSessionEnumerator](https://docs.microsoft.com/en-us/windows/desktop/api/audiopolicy/nn-audiopolicy-iaudiosessionenumerator "IAudioSessionEnumerator")
  - [IAudioSessionControl](https://docs.microsoft.com/en-us/windows/desktop/api/audiopolicy/nn-audiopolicy-iaudiosessioncontrol "IAudioSessionControl")
  - [IAudioSessionControl2](https://docs.microsoft.com/en-us/windows/desktop/api/audiopolicy/nn-audiopolicy-iaudiosessioncontrol2 "IAudioSessionControl2")
  - [ISimpleAudioVolume](https://docs.microsoft.com/en-us/windows/desktop/api/audioclient/nn-audioclient-isimpleaudiovolume "ISimpleAudioVolume")

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
