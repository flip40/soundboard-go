package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	sh "github.com/flip40/soundboard-go/backend/soundhotkey"
	"github.com/gen2brain/malgo"
	"github.com/hajimehoshi/go-mp3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/youpy/go-wav"
)

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

func (b *App) RemoveSound(id string) {
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys = slices.Delete(b.soundHotkeys, i, i+1)
			return
		}
	}
}

func (b *App) PlaySound(path string) {
	deviceInfo, ok := b.playbackDevices[b.selectedDeviceID]
	if !ok {
		b.ErrorDialog("Please select a device first!")
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
