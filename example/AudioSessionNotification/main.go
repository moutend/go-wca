package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-ole/go-ole"

	"github.com/moutend/go-wca/pkg/wca"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type CallbackRegistration struct {
	session        *wca.IAudioSessionControl
	nativeCallback *wca.IAudioSessionEvents
}

func run(args []string) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		return err
	}

	defer ole.CoUninitialize()

	var mmde *wca.IMMDeviceEnumerator
	if err := wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return err
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err := mmde.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
		return err
	}
	defer mmd.Release()

	var ps *wca.IPropertyStore
	if err := mmd.OpenPropertyStore(wca.STGM_READ, &ps); err != nil {
		return err
	}
	defer ps.Release()
	var pv wca.PROPVARIANT
	if err := ps.GetValue(&wca.PKEY_Device_FriendlyName, &pv); err != nil {
		return err
	}
	deviceName := pv.String()
	fmt.Printf("Default playback device: %s\n", deviceName)

	var asm2 *wca.IAudioSessionManager2
	if err := mmd.Activate(wca.IID_IAudioSessionManager2, wca.CLSCTX_ALL, nil, &asm2); err != nil {
		return err
	}
	defer asm2.Release()

	watch := make(chan *wca.IAudioSessionControl, 10)
	release := make(chan CallbackRegistration, 10)

	go func() {
		for session := range watch {
			setupSessionCallback(deviceName, release, session)
		}
	}()

	callback := wca.IAudioSessionNotificationCallback{
		OnSessionCreated: func(pNewSession *wca.IAudioSessionControl) error {
			return onSessionCreated(deviceName, watch, pNewSession)
		},
	}
	asn := wca.NewIAudioSessionNotification(callback)
	if err := asm2.RegisterSessionNotification(asn); err != nil {
		return err
	}

	// You must call IAudioSessionEnumerator::GetCount to begin receiving notifications.
	// https://learn.microsoft.com/en-us/windows/win32/api/audiopolicy/nf-audiopolicy-iaudiosessionmanager2-registersessionnotification
	var sessionEnum *wca.IAudioSessionEnumerator
	if err := asm2.GetSessionEnumerator(&sessionEnum); err != nil {
		return err
	}
	var sessionCount int
	if err := sessionEnum.GetCount(&sessionCount); err != nil {
		return err
	}
	fmt.Printf("%s: %d session(s)\n", deviceName, sessionCount)

	for i := 0; i < sessionCount; i++ {
		var session *wca.IAudioSessionControl
		if err := sessionEnum.GetSession(i, &session); err != nil {
			return err
		}

		session.AddRef()
		watch <- session
	}

	for done := false; !done; {
		select {
		case <-quit:
			fmt.Println("Received keyboard interrupt.")
			done = true

		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			done = true

		case registration := <-release:
			if err := registration.session.UnregisterAudioSessionNotification(registration.nativeCallback); err != nil {
				fmt.Printf("Error unregistering audio session notification: %v\n", err)
			}
			registration.session.Release()
		}
	}

	fmt.Println("Done")

	return nil
}

func onSessionCreated(deviceName string, watch chan *wca.IAudioSessionControl, pNewSession *wca.IAudioSessionControl) error {
	fmt.Printf("%s: called OnSessionCreated\n", deviceName)

	pNewSession.AddRef()
	watch <- pNewSession

	return nil
}

func setupSessionCallback(deviceName string, release chan CallbackRegistration, session *wca.IAudioSessionControl) error {
	var sessionName string
	if err := session.GetDisplayName(&sessionName); err != nil {
		sessionName = fmt.Sprintf("error: %v", err)
	}

	fmt.Printf("%s: session %q\n", deviceName, sessionName)

	var releaseFunc func()

	callback := wca.IAudioSessionEventsCallback{
		OnDisplayNameChanged: func(newDisplayName string, eventContext *ole.GUID) error {
			return onDisplayNameChanged(deviceName, newDisplayName, eventContext)
		},
		OnIconPathChanged: func(newIconPath string, eventContext *ole.GUID) error {
			return onIconPathChanged(deviceName, newIconPath, eventContext)
		},
		// https://github.com/golang/go/issues/45300
		// OnSimpleVolumeChanged: func(newVolume float32, mute bool, eventContext *ole.GUID) error {
		// 	return onSimpleVolumeChanged(deviceName, newVolume, mute, eventContext)
		// },
		// OnChannelVolumeChanged: func(channelCount int, newChannelVolumeArray []float32, changedChannel int, eventContext *ole.GUID) error {
		// 	return onChannelVolumeChanged(deviceName, channelCount, newChannelVolumeArray, changedChannel, eventContext)
		// },
		OnGroupingParamChanged: func(newGroupingParam, eventContext *ole.GUID) error {
			return onGroupingParamChanged(deviceName, newGroupingParam, eventContext)
		},
		OnStateChanged: func(newState wca.AudioSessionState) error {
			return onStateChanged(deviceName, sessionName, releaseFunc, newState)
		},
		OnSessionDisconnected: func(disconnectReason wca.AudioSessionDisconnectReason) error {
			return onSessionDisconnected(deviceName, sessionName, releaseFunc, disconnectReason)
		},
	}
	ase := wca.NewIAudioSessionEvents(callback)
	releaseFunc = func() {
		release <- CallbackRegistration{session, ase}
	}
	err := session.RegisterAudioSessionNotification(ase)
	if err != nil {
		fmt.Printf("Error registering audio session notification: %v\n", err)
	}

	return err
}

func onDisplayNameChanged(deviceName, newDisplayName string, eventContext *ole.GUID) error {
	fmt.Printf("%s: called OnDisplayNameChanged\t%q\n", deviceName, newDisplayName)

	return nil
}

func onIconPathChanged(deviceName, newIconPath string, eventContext *ole.GUID) error {
	fmt.Printf("%s: called OnIconPathChanged\t%q\n", deviceName, newIconPath)

	return nil
}

func onSimpleVolumeChanged(deviceName string, newVolume float32, mute bool, eventContext *ole.GUID) error {
	fmt.Printf("%s: called OnSimpleVolumeChanged\t%f %v\n", deviceName, newVolume, mute)

	return nil
}

func onChannelVolumeChanged(deviceName string, channelCount int, newChannelVolumeArray []float32, changedChannel int, eventContext *ole.GUID) error {
	fmt.Printf("%s: called onChannelVolumeChanged\t%d %v %d\n", deviceName, channelCount, newChannelVolumeArray, changedChannel)

	return nil
}

func onGroupingParamChanged(deviceName string, newGroupingParam, eventContext *ole.GUID) error {
	fmt.Printf("%s: called OnGroupingParamChanged\t%s\n", deviceName, newGroupingParam.String())

	return nil
}

func onStateChanged(deviceName, sessionName string, release func(), newState wca.AudioSessionState) error {
	fmt.Printf("%s: called OnStateChanged %q\t%d\n", deviceName, sessionName, newState)

	if newState == wca.AudioSessionStateExpired {
		release()
	}

	return nil
}

func onSessionDisconnected(deviceName, sessionName string, release func(), disconnectReason wca.AudioSessionDisconnectReason) error {
	fmt.Printf("%s: called OnSessionDisconnected %q\t%d\n", deviceName, sessionName, disconnectReason)
	release()

	return nil
}
