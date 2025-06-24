package app

import (
	sh "github.com/flip40/soundboard-go/backend/soundhotkey"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (b *App) GetSoundHotkeys() []*sh.SoundHotkey {
	return b.soundHotkeys
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

func (b *App) ClearHotkey(id string) {
	for i, sh := range b.soundHotkeys {
		if sh.ID.String() == id {
			b.soundHotkeys[i].Hotkey = nil
		}
	}
}

func (b *App) GetStopHotkey() []uint16 {
	return b.stopHotkey
}

func (b *App) SetStopHotkey(hotkey []uint16) {
	b.stopHotkey = hotkey
}

func (b *App) ClearStopHotkey() {
	b.stopHotkey = nil
}
