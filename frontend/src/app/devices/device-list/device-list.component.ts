import { Component, inject } from '@angular/core';

// TODO: This almost certainly should be passed in by App initially and updated in a different way...
import { GetPlaybackDeviceInfo, SetPlaybackDevice } from "wailsjs/go/main/App"
import { audiodevice } from 'wailsjs/go/models';
import { AudioDeviceService } from 'src/app/shared/audio-device.service';

@Component({
  selector: 'device-list',
  templateUrl: './device-list.component.html',
  styleUrls: ['./device-list.component.scss'],
  standalone: false,
})
export class DeviceListComponent {
  audioDeviceService = inject(AudioDeviceService);
  // @Input() extDevices = [];


  constructor() {
    // TODO: This almost certainly should be passed in by App initially and updated in a different way...
  }

  onChange(audioDevice: audiodevice.AudioDevice) {
    SetPlaybackDevice(this.audioDeviceService.selectedDevice()?.ID ?? "");
  }
}
