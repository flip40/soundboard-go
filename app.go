package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/flip40/soundboard-go/backend/audiodevice"
	sh "github.com/flip40/soundboard-go/backend/soundhotkey"
	"github.com/gen2brain/malgo"
	"github.com/hajimehoshi/go-mp3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/youpy/go-wav"

	hook "github.com/robotn/gohook"
)

type App struct {
	ctx context.Context
	// keypressChannel chan string
	audioCtx *malgo.AllocatedContext

	playbackDevices  map[string]malgo.DeviceInfo
	selectedDeviceID string
	// playbackDeviceID string

	soundHotkeys []*sh.SoundHotkey
	keysDown     map[uint16]struct{}
}

func NewApp() *App {
	return &App{}
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx

	// DEBUG
	b.soundHotkeys = []*sh.SoundHotkey{
		// 	Hotkey: []uint16{162, 87}, // CTRL+W
		// 	Hotkey: []uint16{66}, // B
		sh.NewSoundHotkey("path-to-sound-file", []uint16{66}),
		sh.NewSoundHotkey("path-to-sound-file", []uint16{162, 87}),
		sh.NewSoundHotkey("no hotkey set", nil),
	}

	// AUDIO SETUP
	audioCtx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		panic(err)
	}
	b.audioCtx = audioCtx

	audioDevices, err := audioCtx.Devices(malgo.Playback)
	if err != nil {
		panic(err)
	}
	b.playbackDevices = make(map[string]malgo.DeviceInfo)
	for _, device := range audioDevices {
		b.playbackDevices[device.ID.String()] = device
	}

	// KEY HOOK SETUP
	b.keysDown = make(map[uint16]struct{})
	go b.keyEventProcessor()
}

func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	hook.End()
	// close(b.keypressChannel)
	// portaudio.Terminate()
	_ = b.audioCtx.Uninit()
	b.audioCtx.Free()
	// if b.playbackDevice != nil {
	// 	b.playbackDevice.Uninit()
	// }
}

func (b *App) domReady(ctx context.Context) {
	// b.ShowDialog()
}

func (b *App) ShowDialog(message string) {
	if _, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Native Dialog from Go",
		Message: message,
	}); err != nil {
		panic(err)
	}
}

func (b *App) AddSounds() {
	sounds, err := runtime.OpenMultipleFilesDialog(b.ctx, runtime.OpenDialogOptions{
		// DefaultDirectory: "",
		Title: "Add Sounds...",
		Filters: []runtime.FileFilter{
			{DisplayName: "Sound File", Pattern: "*.wav;*.mp3"},
		},
	})
	if err != nil {
		panic(err)
	}

	// check if no files were returned
	if sounds == nil {
		return
	}

	for _, sound := range sounds {
		b.soundHotkeys = append(b.soundHotkeys, sh.NewSoundHotkey(sound, nil))
	}
}

// func (b *App) GetPlaybackDeviceInfo() []malgo.DeviceInfo {
// 	audioDevices, err := b.audioCtx.Devices(malgo.Playback)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return audioDevices
// }

func (b *App) GetPlaybackDeviceInfo() []audiodevice.AudioDevice {
	// TODO: build this list (and a string map!) on startup
	audioDevices, err := b.audioCtx.Devices(malgo.Playback)
	if err != nil {
		panic(err)
	}

	var wrappedDevices []audiodevice.AudioDevice
	for _, device := range audioDevices {
		wrappedDevices = append(wrappedDevices, audiodevice.AudioDevice{
			ID:   device.ID.String(),
			Name: device.Name(),
		})
	}

	return wrappedDevices
}

func (b *App) SetPlaybackDevice(id string) {
	b.selectedDeviceID = id
}

func (b *App) GetSoundHotkeys() []*sh.SoundHotkey {
	return b.soundHotkeys
}

// func (b *App) SetSoundHotkeys([]sh.SoundHotkey) {

// }

func (b *App) PlaySound(path string) {
	deviceInfo, ok := b.playbackDevices[b.selectedDeviceID]
	if !ok {
		b.ShowDialog("Please select a device first!")
		return
	}

	// TODO: Get path from hotkey input
	// path := "path-to-sound-file"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var reader io.Reader
	var channels, sampleRate uint32
	switch strings.ToLower(filepath.Ext(path)) {
	case ".wav":
		w := wav.NewReader(file)
		f, err := w.Format()
		if err != nil {
			panic(err)
		}

		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

	case ".mp3":
		m, err := mp3.NewDecoder(file)
		if err != nil {
			panic(err)
		}

		reader = m
		channels = 2
		sampleRate = uint32(m.SampleRate())
	default:
		// TODO: Warn with dialogue and return instead of panic
		panic("Not a valid file.")
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.DeviceID = deviceInfo.ID.Pointer()
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = channels
	deviceConfig.SampleRate = sampleRate
	deviceConfig.Alsa.NoMMap = 1

	stopChan := make(chan struct{})
	var stopOnce sync.Once
	stopFunc := func() {
		close(stopChan)
	}
	deviceCallbacks := malgo.DeviceCallbacks{
		Data: func(pOutputSample, pInputSamples []byte, framecount uint32) {
			// TODO: need to make some kind of interrupt here ? Use sync.Once
			if _, err := io.ReadFull(reader, pOutputSample); err != nil {
				stopOnce.Do(stopFunc)
			}
		},
		Stop: func() {
			// TODO: need to make some kind of interrupt here ? Use sync.Once
			// stopChan <- struct{}{}
			stopOnce.Do(stopFunc)
		},
	}

	device, err := malgo.InitDevice(b.audioCtx.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		panic(err)
	}
	defer device.Uninit()

	if err = device.Start(); err != nil {
		panic(err)
	}

	// TODO: add go func to listen for stop key

	var counter int
Playing:
	for {
		select {
		case _, ok := <-stopChan:
			if !ok {
				break Playing
			}
		// TODO: REMOVE DEBUG!!!
		case <-time.After(time.Second):
			counter++
			runtime.LogDebugf(b.ctx, "%d", counter)
		}
	}
}

func (b *App) keyEventProcessor() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		// IGNORE MOUSE EVENTS and IGNORE KEY HOLD
		if ev.Kind >= 6 || ev.Kind == hook.KeyHold {
			continue
		}

		switch ev.Kind {
		case hook.KeyDown:
			b.keysDown[ev.Rawcode] = struct{}{}
		case hook.KeyUp:
			delete(b.keysDown, ev.Rawcode)
		}

		b.checkHotkeys()
	}
}

func (b *App) checkHotkeys() {
	for _, sound := range b.soundHotkeys {
		// skip sounds with no hotkey
		if len(sound.Hotkey) == 0 {
			continue
		}

		// Note: currently disabled to make this easier to trigger rather than harder
		// TODO: make this a toggleable option in settings
		// skip if length of keys pressed doesn't match hotkey
		// if len(keysDown) != len(sound.Hotkey) {
		// 	continue
		// }

		isPressed := true
		for _, key := range sound.Hotkey {
			if _, ok := b.keysDown[key]; !ok {
				isPressed = false
			}
		}

		if isPressed {
			// TODO: PLAY SOUND
			// fmt.Printf("Playing Sound for '%s'\n", sound.Path)
			b.PlaySound(sound.Path)
		}
	}
}

func (b *App) SetHotkey(id string, hotkey []uint16) {
	fmt.Printf("\n\nID TEST %+v\n\n", id)
	// idAsUUID, _ := uuid.FromBytes(id[:])
	fmt.Printf("setting hotkey %s to %+v\n", id, hotkey)
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys[i].Hotkey = hotkey
		}
	}
	// TODO: log error if not found?

	// TODO: DEBUG

	fmt.Printf("Hotkeys after set: %+v\n", b.soundHotkeys)
}

// b.soundHotkeys = slices.Delete(b.soundHotkeys, i, i+1)
// return

func (b *App) ClearHotkey(id string) {
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys[i].Hotkey = nil
		}
	}

	// TODO: DEBUG
	fmt.Printf("Hotkeys after clear: %+v\n", b.soundHotkeys)
}
