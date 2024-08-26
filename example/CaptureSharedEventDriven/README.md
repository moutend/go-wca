# Capturing audio with shared event driven mode

This example shows that the capturing audio with shared event driven mode.

## Prerequisites

- Go 1.13 or later

## Build

```console
go build
```

That's it. Then you'll get `CaptureSharedEventDriven.exe`. Note that your platform is not Windows, you need set the environment variable `GOOS='windows'`.

## Usage

```console
./CaptureSharedEventDriven -o music.wav -d 10
```

Please specify the flag `-o` or `--output` for saving audio file. The `-d` or `--duration` is optional and it indicates recording duration in second. If the recording duration was not specified, it keeps recording until receiving interruption by Ctrl-C.

## Note

Stability of the event driven mode is much lower than the timer driven mode. Because we cannot control the scheduling of goroutines, when the goroutine which observing audio ready event was stopped, stutterring occurs.

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

The Windows Core Audio API was introduced Windows vista, so that the later than that version of Windows could run this example. However, I'm not sure because I've just tested this example on Windows 10 version 1607 at the moment. Operation verification including bug report are welcome.
