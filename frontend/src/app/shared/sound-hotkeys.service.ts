import { Injectable, signal } from '@angular/core';
import { GetSoundHotkeys, GetStopHotkey, RemoveSound } from 'wailsjs/go/main/App';
import { soundhotkey } from 'wailsjs/go/models';

@Injectable({
  providedIn: 'root',
})
export class SoundHotkeysService {
  soundHotkeys = signal<soundhotkey.SoundHotkey[]>([]);
  stopHotkey = signal<number[] | undefined>(undefined);

  constructor() {
    this.updateHotkeys();
  }

  getHotkeys(): soundhotkey.SoundHotkey[] {
    return this.soundHotkeys();
  }

  getHotkeyByID(id: number[]): soundhotkey.SoundHotkey | undefined {
    return this.soundHotkeys().find(soundHotkey => soundHotkey.ID == id);
  }

  getStopHotkey(): number[] | undefined {
    return this.stopHotkey();
  }

  updateHotkeys() {
    GetSoundHotkeys().then(soundHotkeys => this.soundHotkeys.set(soundHotkeys));
    GetStopHotkey().then(stopHotkey => this.stopHotkey.set(stopHotkey));
  }

  removeSound(hotkeyID: number[]) {
    RemoveSound(String(hotkeyID)).then(() => this.updateHotkeys());
  }
}
