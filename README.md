# go-wca (beta)

[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![GitHub release](http://img.shields.io/github/release/moutend/go-wca.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[godocs]: http://godoc.org/github.com/moutend/go-wca
[release]: https://github.com/moutend/go-wca/releases
[license]: https://github.com/moutend/go-wca/blob/master/LICENSE

Go bindings for Windows Core Audio API without using cgo.

This package allows you to do:

- Rendering audio with shared timer driven mode
- Rendering audio with event driven mode
- Capturing audio with shared timer driven mode
- Capturing audio with event driven mode
- Loopback capturing with shared timer mode
- Loopback capturing with shared event mode

If you're not familiar with Windows Core Audio API, [the documentation on MSDN](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370802(v=vs.85).aspx) helps you to learn about shared / exclusive mode audio.

## Prerequisites

- Go 1.8.1 or later
- `go-ole` ([github.com/go-ole/go-ole](https://github.com/go-ole/go-ole))

## Examples

You can find the examples in `example` directory. For more information, please read the README.md in each examples.

## Documentation

Each APIs in this package correspond to native COM APIs, so that you can refer the documentation on MSDN as a full documentation of this package.

The following list contains the links to the documentation of native API which is available in this package.
- MMDevice API
  - [IMMDevice](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371395(v=vs.85).aspx)
  - [IMMDeviceCollection](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371396(v=vs.85).aspx "IMMDeviceCollection"){#mt779369_VS.85_en-us}
  - [IMMDeviceEnumerator](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371399(v=vs.85).aspx "IMMDeviceEnumerator"){#mt779370_VS.85_en-us}
  - [IMMEndpoint](https://msdn.microsoft.com/en-us/library/windows/desktop/dd371414(v=vs.85).aspx "IMMEndpoint"){#mt779371_VS.85_en-us}
- Windows Audio Session API
  - [IAudioClient](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370865(v=vs.85).aspx "IAudioClient"){#mt779327_VS.85_en-us}- [IAudioClient2](https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx "IAudioClient2"){#mt779328_VS.85_en-us}
  - [IAudioClient2](https://msdn.microsoft.com/en-us/library/windows/desktop/hh404179(v=vs.85).aspx "IAudioClient2"){#mt779328_VS.85_en-us}
  - [IAudioClient3](https://msdn.microsoft.com/en-us/library/windows/desktop/dn911487(v=vs.85).aspx "IAudioClient3"){#mt779329_VS.85_en-us}
  - [IAudioCaptureClient](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370858(v=vs.85).aspx "IAudioCaptureClient"){#mt779325_VS.85_en-us}
  - [IAudioEndpointVolume](https://msdn.microsoft.com/en-us/library/windows/desktop/dd370892(v=vs.85).aspx "IAudioEndpointVolume"){#mt779337_VS.85_en-us}
  - [IAudioRenderClient](https://msdn.microsoft.com/en-us/library/windows/desktop/dd368242(v=vs.85).aspx "IAudioRenderClient"){#mt779347_VS.85_en-us}

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
