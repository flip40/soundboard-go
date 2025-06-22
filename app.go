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
		runtime.LogDebugf(b.ctx, "MALGO > %s\n", message)
	})
	if err != nil {
		runtime.LogErrorf(b.ctx, "failed to initialize malgo context: %s", err.Error())
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
		runtime.LogError(b.ctx, err.Error())
	}
}

func (b *App) ErrorDialog(message string) {
	if _, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Error!",
		Message: message,
	}); err != nil {
		runtime.LogError(b.ctx, err.Error())
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
		b.ErrorDialog("Failed to open sound files")
		runtime.LogError(b.ctx, err.Error())
	}

	// check if no files were returned (user cancelled)
	if sounds == nil {
		return
	}

	for _, sound := range sounds {
		b.soundHotkeys = append(b.soundHotkeys, sh.NewSoundHotkey(sound, nil))
	}
}

func (b *App) GetPlaybackDeviceInfo() []audiodevice.AudioDevice {
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

	file, err := os.Open(path)
	if err != nil {
		b.ErrorDialog(fmt.Sprintf("Could not open file at: %s", path))
		runtime.LogError(b.ctx, err.Error())
	}
	defer file.Close()

	var reader io.Reader
	var channels, sampleRate uint32
	switch strings.ToLower(filepath.Ext(path)) {
	case ".wav":
		w := wav.NewReader(file)
		f, err := w.Format()
		if err != nil {
			b.ErrorDialog(fmt.Sprintf("Could not read .wav file at: %s", path))
			runtime.LogError(b.ctx, err.Error())
		}

		reader = w
		channels = uint32(f.NumChannels)
		sampleRate = f.SampleRate

	case ".mp3":
		m, err := mp3.NewDecoder(file)
		if err != nil {
			b.ErrorDialog(fmt.Sprintf("Could not read .mp3 file at: %s", path))
			runtime.LogError(b.ctx, err.Error())
		}

		reader = m
		channels = 2
		sampleRate = uint32(m.SampleRate())
	default:
		b.ErrorDialog(fmt.Sprintf("Could not play file at: %s", path))
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.DeviceID = deviceInfo.ID.Pointer()
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = channels
	deviceConfig.SampleRate = sampleRate
	deviceConfig.Alsa.NoMMap = 1

	stopChan := make(chan struct{})
	var stopOnce sync.Once

	b.stopFuncs[&stopOnce] = func() {
		close(stopChan)
	}
	deviceCallbacks := malgo.DeviceCallbacks{
		Data: func(pOutputSample, pInputSamples []byte, framecount uint32) {
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
		runtime.LogErrorf(b.ctx, "failed to initialize malgo device: %s", err.Error())
		panic(err)
	}
	defer device.Uninit()

	if err = device.Start(); err != nil {
		runtime.LogErrorf(b.ctx, "failed to start playback device: %s", err.Error())
		panic(err)
	}

	// block until stopped
	<-stopChan

	// cleanup stop funcs
	b.stopFuncsMutex.Lock()
	defer b.stopFuncsMutex.Unlock()

	delete(b.stopFuncs, &stopOnce)
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
	found := false
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys[i].Hotkey = hotkey
			found = true
		}
	}

	if !found {
		runtime.LogErrorf(b.ctx, "failed to update hotkey, id wasn't found in sound hotkey map: %s", id)
	}
}

// b.soundHotkeys = slices.Delete(b.soundHotkeys, i, i+1)
// return

func (b *App) ClearHotkey(id string) {
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys[i].Hotkey = nil
		}
	}
}

func (b *App) ClearStopHotkey() {
	b.stopHotkey = nil
}

func (b *App) SetStopHotkey(hotkey []uint16) {
	b.stopHotkey = hotkey
}

func (b *App) RemoveSound(id string) {
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
		b.ErrorDialog("Failed to open soundboard file.")
		runtime.LogError(b.ctx, err.Error())
	}

	// check if no files were returned (user cancelled)
	if sbFilePath == "" {
		return
	}

	sbFileBytes, err := os.ReadFile(sbFilePath)
	if err != nil {
		b.ErrorDialog("Could not read soundboard file.")
		runtime.LogError(b.ctx, err.Error())
	}

	sb := &soundboard.Soundboard{}
	if err := json.Unmarshal(sbFileBytes, sb); err != nil {
		b.ErrorDialog("Invalid soundboard file!")
		runtime.LogError(b.ctx, err.Error())
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
		b.ErrorDialog("Failed to save soundboard file.")
		runtime.LogError(b.ctx, err.Error())
	}

	sb := &soundboard.Soundboard{
		SelectedDeviceID: b.selectedDeviceID,
		SoundHotkeys:     b.soundHotkeys,
		StopHotkey:       b.stopHotkey,
	}

	sbBytes, err := json.Marshal(sb)
	if err != nil {
		b.ErrorDialog("Failed to open soundboard file.")
		runtime.LogError(b.ctx, err.Error())
	}

	if err := os.WriteFile(sbFilePath, sbBytes, 0644); err != nil {
		b.ErrorDialog(fmt.Sprintf("Failed to write soundboard file to %s", sbFilePath))
		runtime.LogError(b.ctx, err.Error())
	}
}
