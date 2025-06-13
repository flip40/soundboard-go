import { Injectable, signal } from '@angular/core';
import { GetSoundHotkeys } from 'wailsjs/go/main/App';
import { soundhotkey } from 'wailsjs/go/models';

@Injectable({
  providedIn: 'root',
})
export class SoundHotkeysService {
  soundHotkeys = signal<soundhotkey.SoundHotkey[]>([]);

  constructor() {
    this.updateHotkeys();
  }

  getHotkeys(): soundhotkey.SoundHotkey[] {
    return this.soundHotkeys();
  }

  getHotkeyByID(id: number[]): soundhotkey.SoundHotkey | undefined {
    return this.soundHotkeys().find(soundHotkey => soundHotkey.ID == id);
  }

  updateHotkeys() {
    GetSoundHotkeys().then(soundHotkeys => this.soundHotkeys.set(soundHotkeys));
  }
}
