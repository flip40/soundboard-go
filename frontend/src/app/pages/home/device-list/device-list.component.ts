import { Component, inject } from '@angular/core';
import { AudioDeviceService } from 'src/app/shared/audio-device.service';
import { SetPlaybackDevice } from "wailsjs/go/app/App"
import { audiodevice } from 'wailsjs/go/models';

@Component({
  selector: 'device-list',
  templateUrl: './device-list.component.html',
  styleUrls: ['./device-list.component.scss'],
  standalone: false,
})
export class DeviceListComponent {
  audioDeviceService = inject(AudioDeviceService);

  constructor() { }

  onChange(audioDevice: audiodevice.AudioDevice) {
    SetPlaybackDevice(this.audioDeviceService.selectedDevice()?.ID ?? "");
  }
}
