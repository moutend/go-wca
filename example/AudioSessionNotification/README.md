# WASAPI Audio Session Notification

An example to capture WASAPI audio session events.

## Prerequisites

- Go 1.13 or later
- `go-ole` (https://github.com/go-ole/go-ole)

## Build

```console
go build
```

That's it. Then you'll get `AudioSessionNotification.exe`. Note that your platform is not Windows, you need set the environment variable `GOOS='windows'`.

## Usage

```console
./AudioSessionNotification
```

## Contributing

Bug reports and improving the documentation are welcome. (https://github.com/moutend/go-wca/issues)

The Windows Core Audio API IAudioSessionNotification was introduced in Windows 7, so that the later than that version of Windows should run this example.
Tested on Windows 10 however.
