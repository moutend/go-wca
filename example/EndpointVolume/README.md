# Endpoint volume

This example shows that changing volume for default rendering device.

## Prerequisites

- Go 1.8.1 or later
- `go-ole` (https://github.com/go-ole/go-ole)

## Build the executable

```shell
go build
```

That's it. Then you'll get `EndpointVolume.exe`. Note that your platform is not Windows, you need set the environment variable `GOOS='windows'` before building the executable.

## Usage

```shell
./EndpointVolume --volume 0.1
```

Available flags are:

- `-v` or `--volume` sets the volume as scalar value
- `-g` or `--gain` sets the volume as level (dB) value
- `-m` or `--mute` sets mute state

## Contributing

Bug reports and improving the documentation are welcome. (https://github.com/moutend/go-wca/issues)

The Windows Core Audio API was introduced Windows vista, so that the later than that version of Windows could run this example. However, I'm not sure because I've just tested this example on Windows 10 version 1607 at the moment.
