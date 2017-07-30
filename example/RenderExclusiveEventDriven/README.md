# Rendering with exclusive event driven mode

This example shows that the rendering audio with exclusive event driven mode.

## Prerequisites

- Go 1.8.1 or later
- `go-ole` (https://github.com/go-ole/go-ole)

## Download

You can download the executable from [Releases page](https://github.com/moutend/go-wca/releases).

## Build

```shell
go build
```

That's it. Then you'll get `RenderExclusiveEventDriven.exe`. Note that your platform is not Windows, you need set the environment variable `GOOS='windows'` before building the executable.

## Usage

```shell
./RenderExclusiveEventDriven -i music.wav
```

Please specify the WAVE audio file with `-i` or `--input` flag.

## Note

Whether the wAV file can play or not is up to hardware settings. You need specify the WAV file which is encoded as same as output settings.

For example, when the output is configured to comibination of 96 kHz and 24 bit, you can only play WAV files which are encoded with 96 kHz and 24 bit.

## Known issues

Some audio interfaces fails playback in exclusive event driven mode. I use Steinberg UR22 as an external audio interface, and the interface can play correctly in timer driven mode, but it makes glitch in event driven mode.

I don't know why but if you want to use exclusive audio rendering, you should use timer driven mode for now.

## Contributing

1. Fork ([https://github.com/moutend/go-wca/fork](https://github.com/moutend/go-wca/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

The Windows Core Audio API was introduced Windows vista, so that the later than that version of Windows could run this example. However, I'm not sure because I've just tested this example on Windows 10 version 1607 at the moment. Operation verification including bug report are welcome.
