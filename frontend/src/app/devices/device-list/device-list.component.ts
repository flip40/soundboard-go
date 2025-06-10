import { Component } from '@angular/core';

// TODO: This almost certainly should be passed in by App initially and updated in a different way...
import { GetPlaybackDeviceInfo, SetPlaybackDevice } from "wailsjs/go/main/App"
import { audiodevice } from 'wailsjs/go/models';

@Component({
  selector: 'device-list',
  templateUrl: './device-list.component.html',
  styleUrls: ['./device-list.component.scss'],
  standalone: false,
})
export class DeviceListComponent {
  // @Input() extDevices = [];

  devices: audiodevice.AudioDevice[] = [];
  selectedDevice: string = "";

  constructor() {
    // TODO: This almost certainly should be passed in by App initially and updated in a different way...
    GetPlaybackDeviceInfo().then((devices) => {
      this.devices = devices;
    });
  }

  onChange(event: Event) {
    SetPlaybackDevice(this.selectedDevice);
  }
}
