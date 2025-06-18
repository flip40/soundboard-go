// Package soundboard represents a save file for a complete soundboard
package soundboard

import (
	sh "github.com/flip40/soundboard-go/backend/soundhotkey"
)

type Soundboard struct {
	SelectedDeviceID string            `json:"selected_device"`
	SoundHotkeys     []*sh.SoundHotkey `json:"sound_hotkeys"`
	StopHotkey       sh.Hotkey         `json:"stop_hotkey"`
}
