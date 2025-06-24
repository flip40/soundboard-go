package soundhotkey

import (
	"fmt"
	"strings"

	"github.com/flip40/soundboard-go/backend/keycodes"
	"github.com/google/uuid"
)

type Hotkey []uint16

func (hotkey Hotkey) String() string {
	keys := []string{}
	for _, keycode := range hotkey {
		keys = append(keys, keycodes.RawcodeToString(keycode))
	}
	return strings.Join(keys, "+")
}

type SoundHotkey struct {
	ID     uuid.UUID `json:"id"`
	Path   string    `json:"path"`
	Hotkey Hotkey    `json:"hotkey"`
}

func NewSoundHotkey(path string, hotkey []uint16) *SoundHotkey {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return &SoundHotkey{
		ID:     uuid,
		Path:   path,
		Hotkey: hotkey,
	}
}

func (sh *SoundHotkey) String() string {
	return fmt.Sprintf("%+v", *sh)
}
