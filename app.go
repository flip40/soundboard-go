package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/flip40/soundboard-go/backend/audiodevice"
	"github.com/flip40/soundboard-go/backend/soundboard"
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
	stopHotkey   sh.Hotkey
	keysDown     map[uint16]struct{}

	stopFuncs      map[*sync.Once]func()
	stopFuncsMutex sync.Mutex
}

func NewApp() *App {
	return &App{}
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
	b.stopFuncs = make(map[*sync.Once]func())

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
		// TODO: Error?
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

	var wrappedDevices audiodevice.AudioDevices
	for _, device := range audioDevices {
		wrappedDevices = append(wrappedDevices, audiodevice.AudioDevice{
			ID:       device.ID.String(),
			Name:     device.Name(),
			Selected: func() bool { return device.ID.String() == b.selectedDeviceID }(),
		})
	}

	sort.Sort(wrappedDevices)

	return wrappedDevices
}

func (b *App) SetPlaybackDevice(id string) {
	b.selectedDeviceID = id
}

func (b *App) GetSoundHotkeys() []*sh.SoundHotkey {
	return b.soundHotkeys
}

func (b *App) GetStopHotkey() []uint16 {
	return b.stopHotkey
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
	// stopFunc := func() {
	// 	close(stopChan)
	// }
	b.stopFuncs[&stopOnce] = func() {
		close(stopChan)
	}
	deviceCallbacks := malgo.DeviceCallbacks{
		Data: func(pOutputSample, pInputSamples []byte, framecount uint32) {
			// TODO: need to make some kind of interrupt here ? Use sync.Once
			if _, err := io.ReadFull(reader, pOutputSample); err != nil {
				stopOnce.Do(b.stopFuncs[&stopOnce])
			}
		},
		Stop: func() {
			stopOnce.Do(b.stopFuncs[&stopOnce])
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
				b.stopFuncsMutex.Lock()
				defer b.stopFuncsMutex.Unlock()

				delete(b.stopFuncs, &stopOnce)
				break Playing
			}
		// TODO: REMOVE DEBUG!!!
		case <-time.After(time.Second):
			counter++
			runtime.LogDebugf(b.ctx, "%d", counter)
		}
	}
}

func (b *App) StopAllSounds() {
	b.stopFuncsMutex.Lock()
	defer b.stopFuncsMutex.Unlock()

	for stopOnce, stopFunc := range b.stopFuncs {
		stopOnce.Do(stopFunc)
		delete(b.stopFuncs, stopOnce)
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
		b.checkStopHotkey()
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
				break
			}
		}

		if isPressed {
			// TODO: PLAY SOUND
			// fmt.Printf("Playing Sound for '%s'\n", sound.Path)
			go b.PlaySound(sound.Path)
		}
	}
}

func (b *App) checkStopHotkey() {
	// skip if stop hotkey is not set
	if len(b.stopHotkey) == 0 {
		return
	}

	isPressed := true
	for _, key := range b.stopHotkey {
		if _, ok := b.keysDown[key]; !ok {
			isPressed = false
			break
		}
	}

	if isPressed {
		go b.StopAllSounds()
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

func (b *App) ClearStopHotkey() {
	b.stopHotkey = nil
}

func (b *App) SetStopHotkey(hotkey []uint16) {
	b.stopHotkey = hotkey
}

func (b *App) RemoveSound(id string) {
	fmt.Printf("removing sound: %+v\n", id)
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys = slices.Delete(b.soundHotkeys, i, i+1)
			return
		}
	}
}
func (b *App) ResetSoundboard() {
	b.StopAllSounds()

	b.selectedDeviceID = ""
	b.soundHotkeys = []*sh.SoundHotkey{}
	b.stopHotkey = nil
}

func (b *App) LoadSoundboard() {
	b.StopAllSounds()

	sbFilePath, err := runtime.OpenFileDialog(b.ctx, runtime.OpenDialogOptions{
		// DefaultDirectory: "",
		Title: "Load Soundboard...",
		Filters: []runtime.FileFilter{
			{DisplayName: "Soundboard File", Pattern: "*.sb;*.json"},
		},
	})
	if err != nil {
		panic(err)
	}

	// check if no files were returned
	if sbFilePath == "" {
		// TODO: Error?
		return
	}

	sbFileBytes, err := os.ReadFile(sbFilePath)
	if err != nil {
		panic(err)
	}

	sb := &soundboard.Soundboard{}
	if err := json.Unmarshal(sbFileBytes, sb); err != nil {
		// TODO: handle error message
		panic(err)
	}

	b.selectedDeviceID = sb.SelectedDeviceID
	b.soundHotkeys = sb.SoundHotkeys
	b.stopHotkey = sb.StopHotkey
}

func (b *App) SaveSoundboard() {
	sbFilePath, err := runtime.SaveFileDialog(b.ctx, runtime.SaveDialogOptions{
		// DefaultDirectory: "",
		Title: "Load Soundboard...",
		Filters: []runtime.FileFilter{
			{DisplayName: "Soundboard File", Pattern: "*.sb;*.json"},
		},
	})
	if err != nil {
		panic(err)
	}

	sb := &soundboard.Soundboard{
		SelectedDeviceID: b.selectedDeviceID,
		SoundHotkeys:     b.soundHotkeys,
		StopHotkey:       b.stopHotkey,
	}

	sbBytes, err := json.Marshal(sb)
	if err != nil {
		// TODO: display error
		panic(err)
	}

	if err := os.WriteFile(sbFilePath, sbBytes, 0644); err != nil {
		// TODO: display error
		panic(err)
	}
}
