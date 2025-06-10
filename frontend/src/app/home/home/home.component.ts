import { Component, Input } from '@angular/core';
// import { GetKeycodes } from 'wailsjs/go/keycodes/KeycodeHelper';

@Component({
    selector: 'home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
    standalone: false,
})
export class HomeComponent {
  // @Input() keycodes: Record<number, string> = {};

  constructor() {

    // GetKeycodes().then((keycodes) => {
    //   this.keycodes = keycodes;
    // });
  }


      // GetPlaybackDeviceInfo().then((devices) => {
      //   this.devices = devices;
      // });
}
