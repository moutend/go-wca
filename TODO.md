# TODO

## Use go.mod

There are no reasons to not use go.mod.

## Change project layout

Follow the standard golang project layout. The project root directory contains many .go codes and it looks dirty.

https://github.com/golang-standards/project-layout

## Fix CI

Setup CI to use latest golang.

## Update README and LICENSE

e.g. Copyright year.

## Add tests

I test the example executables on my Windows 10 laptop (more precisely, bootcamp with Macbook Air) at the moment.
I'm looking for CI which targets Windows client apps. For example, Appveyer CI supports only Windows Server and the server edition doesn't support physical audio device.

## Figure out how to handle bizarre default bit depth

`IAudioClient::GetMixFormat` returns always 32 bit as a bit depth on my machine (Macbook Air / Windows 10 version 1607).
I'm investigating this is my machine specific issue or not.
FYI, I don't know why, but The sample rate seems to be always correct (e.g. 44100, 48000 and so on).
