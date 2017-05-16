# Loopback capturing with shared event driven mode

This example shows that the loopback capturing with shared event driven mode.

Capturing loopback audio with event driven mode has an issue.

## Prerequisites

- Go 1.8.1 or later
- `go-ole` (https://github.com/go-ole/go-ole)

## Build the executable

```shell
go build
```

That's it. Then you'll get `LoopbackCaptureSharedEventDriven.exe`. Note that your platform is not Windows, you need set the environment variable `GOOS='windows'` before building the executable.

## Usage

```shell
./LoopbackCaptureSharedEventDriven -o music.wav -d 10
```

Please specify the flag `-o` or `--output` for saving audio file. The `-d` or `--duration` is optional and it indicates recording duration in second. If the recording duration was not specified, it keeps recording until receiving interruption by Ctrl-C.

## Issue that the audio ready event never fire

According to the [MSDN's documentation about loopback capturing](https://msdn.microsoft.com/en-us/library/windows/desktop/dd316551(v=vs.85).aspx), the audio capture client can set the audio ready event without any errors, but it will never fire.

> A pull-mode capture client does not receive any events when a stream is initialized with event-driven buffering and is loopback-enabled.
> To work around this, initialize a render stream in event-driven mode. Each time the client receives an event for the render stream, it must signal the capture client to run the capture thread that reads the next set of samples from the capture endpoint buffer.
>
> https://msdn.microsoft.com/en-us/library/windows/desktop/dd316551(v=vs.85).aspx

I don't know why this issue has not fixed yet, but we need add the workaround for current implementation of Core Audio API.

In this example, I initializes two audio clients which correspond to capturing and rendering (main.go:170 and main.go:175) as a workaround. It seems to work fine for now, but I think that the timer driven mode is stable and proper way to capture loopback audio.

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
