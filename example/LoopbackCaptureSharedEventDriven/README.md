# Loopback capturing with shared event driven mode

This example shows that the loopback (a.k.a. what you hear) capturing with shared event driven mode.

Capturing loopback audio with event driven mode has an issue.

## Prerequisites

- Go 1.8 or later
- `go-ole` (https://github.com/go-ole/go-ole)

## Download

You can download the executable from [Releases page](https://github.com/moutend/go-wca/releases).

## Build

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
> 
> To work around this, initialize a render stream in event-driven mode. Each time the client receives an event for the render stream, it must signal the capture client to run the capture thread that reads the next set of samples from the capture endpoint buffer.

I don't know why this issue has not been fixed yet, but we need add the workaround for current implementation of Windows Core Audio API.

## Note

As with capturing from microphone, stability of the event driven mode is much lower than the timer driven mode. Because we cannot control the scheduling of goroutines, when the goroutine which observing audio ready event was stopped, stutterring occurs.

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

The Windows Core Audio API was introduced Windows vista, so that the later than that version of Windows could run this example. However, I'm not sure because I've just tested this example on Windows 10 version 1607 at the moment. Operation verification including bug report are welcome.
