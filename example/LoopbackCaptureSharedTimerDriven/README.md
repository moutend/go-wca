# Loopback capturing with shared timer driven mode

This example shows that the loopback capturing with shared timer driven mode.

## Prerequisites

- Go 1.8.1 or later
- `go-ole` (https://github.com/go-ole/go-ole)

## Build the executable

```shell
go build
```

That's it. Then you'll get `LoopbackCaptureSharedTimerDriven.exe`. Note that your platform is not Windows, you need set the environment variable `GOOS='windows'` before building the executable.

## Usage

```shell
./LoopbackCaptureSharedTimerDriven -o music.wav -d 10
```

Please specify the flag `-o` or `--output` for saving audio file. The `-d` or `--duration` is optional and it indicates recording duration in second. If the recording duration was not specified, it keeps recording until receiving interruption by Ctrl-C.

## Note

This example captures the audio as 44100 Hz / 16 bit wave format audio. The capturing step will be failed when the sample rate and bit depth of system default **playback** device is not set as that value.

To avoid this error, please set the system default settings for shared mode recording.

1. Open control panel and select sound.
1. Select playback tab and open property of the main playback device.
1. Select the advanced tab, set the default sample rate and bit depth as 44100 Hz / 16 bit.
1. Apply changes.

## Contributing

Bug reports and improving the documentation are welcome. (https://github.com/moutend/go-wca/issues)

The Windows Core Audio API was introduced Windows vista, so that the later than that version of Windows could run this example. However, I'm not sure because I've just tested this example on Windows 10 version 1607 at the moment.
