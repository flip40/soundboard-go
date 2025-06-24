package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/flip40/soundboard-go/backend/soundboard"
	sh "github.com/flip40/soundboard-go/backend/soundhotkey"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (b *App) ResetSoundboard() {
	b.StopAllSounds()

	b.selectedDeviceID = ""
	b.soundHotkeys = []*sh.SoundHotkey{}
	b.stopHotkey = nil
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
