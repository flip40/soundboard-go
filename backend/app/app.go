package app

import (
	"context"
	"sync"

	"github.com/flip40/soundboard-go/backend/audiodevice"
	sh "github.com/flip40/soundboard-go/backend/soundhotkey"
	"github.com/gen2brain/malgo"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	hook "github.com/robotn/gohook"
)

// TODO: a lot of functions here can be more neatly organized into separate files
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

type Exported interface {
	ErrorDialog(message string)

	GetPlaybackDeviceInfo() []audiodevice.AudioDevice
	SetPlaybackDevice(id string)

	AddSounds()
	RemoveSound(id string)
	PlaySound(path string)
	StopAllSounds()

	GetSoundHotkeys() []*sh.SoundHotkey
	SetHotkey(id string, hotkey []uint16)
	ClearHotkey(id string)

	GetStopHotkey() []uint16
	SetStopHotkey(hotkey []uint16)
	ClearStopHotkey()

	ResetSoundboard()
	LoadSoundboard()
	SaveSoundboard()
}

func NewApp() *App {
	return &App{}
}

func (b *App) Startup(ctx context.Context) {
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

func (b *App) Shutdown(ctx context.Context) {
	// stop key hook
	hook.End()

	// free audio context
	_ = b.audioCtx.Uninit()
	b.audioCtx.Free()
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
