package app

import (
	hook "github.com/robotn/gohook"
)

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

// check if any hotkeys are pressed in current key state
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
