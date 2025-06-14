import { Injectable, signal } from '@angular/core';
import { GetSoundHotkeys } from 'wailsjs/go/main/App';
import { soundhotkey } from 'wailsjs/go/models';
import { GetPlaybackDeviceInfo, SetPlaybackDevice } from "wailsjs/go/main/App"
import { audiodevice } from 'wailsjs/go/models';

@Injectable({
  providedIn: 'root',
})
export class AudioDeviceService {
  selectedDevice = signal<audiodevice.AudioDevice | undefined>(undefined);
  audioDevices = signal<audiodevice.AudioDevice[]>([]);

  constructor() {
    this.updateAudioDevices();
  }

  updateAudioDevices() {
    GetPlaybackDeviceInfo().then((audioDevices) => {
      this.audioDevices.set(audioDevices);
    });
  }
}
