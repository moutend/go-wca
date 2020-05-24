# TODO(s)

## Test audio devices on the CI server

I test the `_example` on my Windows 10 laptop at the moment. I'm looking for CI which can test the audio device. For example, Appveyer CI supports only Windows Server and the server edition doesn't support physical audio device.

## Figure out how to handle bizarre default bit depth

`IAudioClient::GetMixFormat` returns always 32 bit as a bit depth on my machine (Macbook Air / Windows 10 version 1607). I'm investigating this is my machine specific issue or not.
