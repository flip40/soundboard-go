import { Injectable, signal } from '@angular/core';
import { GetPlaybackDeviceInfo } from "wailsjs/go/app/App"
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
      this.selectedDevice.set(audioDevices.find(audioDevice => audioDevice.Selected));
    });
  }
}
