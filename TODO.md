# Figure out how to handle IMMNotificationClient

iMMNotificationClient is a special interface and we should implement our own.

So, I tried implementing IMMNotificationClient with `syscall.NewCallback`. I did run the code below, wait a few seconds, and then plug in USB headset, but the output was nothing.

```go
// +build windows
package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca"
)

func main() {
	var err error
	if err = run(os.Args); err != nil {
		log.Fatal(err)
	}
	return
}

func run(args []string) (err error) {
	return foo()
}

func foo() (err error) {
	if err = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return
	}
	defer ole.CoUninitialize()

	var mmde *wca.IMMDeviceEnumerator
	if err = wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err = mmde.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
		return
	}
	defer mmd.Release()

	var ps *wca.IPropertyStore
	if err = mmd.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		return
	}
	defer ps.Release()

	var pv wca.PROPVARIANT
	if err = ps.GetValue(&wca.PKEY_Device_FriendlyName, &pv); err != nil {
		return
	}
	fmt.Printf("%s\n", pv.String())

	mmnc := &wca.IMMNotificationClient{}
	mmnc.VTable = &wca.IMMNotificationClientVtbl{}
	mmnc.VTable.QueryInterface = syscall.NewCallback(queryInterface)
	mmnc.VTable.AddRef = syscall.NewCallback(addRef)
	mmnc.VTable.Release = syscall.NewCallback(release)
	mmnc.VTable.OnDeviceStateChanged = syscall.NewCallback(OnDeviceStateChanged)
	mmnc.VTable.OnDeviceAdded = syscall.NewCallback(OnDeviceAdded)
	mmnc.VTable.OnDeviceRemoved = syscall.NewCallback(OnDeviceRemoved)
	mmnc.VTable.OnDefaultDeviceChanged = syscall.NewCallback(OnDefaultDeviceChanged)
	mmnc.VTable.OnPropertyValueChanged = syscall.NewCallback(OnPropertyValueChanged)

	if err = mmde.RegisterEndpointNotificationCallback(mmnc); err != nil {
		return
	}
	time.Sleep(10 * time.Second)
	return
}

func queryInterface(this *wca.IMMNotificationClient, iid *ole.GUID, punk **ole.IUnknown) (hresult uintptr) {
	fmt.Println("called")
	return
}
func addRef(this *wca.IMMNotificationClient) (hResult uintptr) {
	fmt.Println("called addref")
	return
}
func release(this *wca.IMMNotificationClient) (hResult uintptr) {
	fmt.Println("called release")
	return
}
func OnDeviceStateChanged(this *wca.IMMNotificationClient, lctwstr uintptr, state uint32) (hResult uintptr) {
	fmt.Println("called OnDeviceStateChanged")
	return
}
func OnDeviceAdded(this *wca.IMMNotificationClient, lpcwstr uintptr) (hResult uintptr) {
	fmt.Println("called OnDeviceAdded")
	return
}
func OnDeviceRemoved(this *wca.IMMNotificationClient, lpcwstr uintptr) (hResult uintptr) {
	fmt.Println("called OnDeviceRemoved")
	return
}
func OnDefaultDeviceChanged(this *wca.IMMNotificationClient, EDataFlow, eRole uint32, lpcwstr uintptr) (hResult uintptr) {
	fmt.Println("called OnDefaultDeviceChanged")
	return
}
func OnPropertyValueChanged(this *wca.IMMNotificationClient, lpcwstr uintptr, key uint32) (hResult uintptr) {
	fmt.Println("called OnPropertyValueChanged")
	return
}
```

# Figure out how to handle bizarre default bit depth

`IAudioClient::GetMixFormat` returns always 32 bit as a bit depth on my machine (Macbook Air / Windows 10 version 1607).
I'm investigating this is my machine specific issue or not.
FYI, I don't know why, but The sample rate seems to be always correct (e.g. 44100, 48000 and so on).
