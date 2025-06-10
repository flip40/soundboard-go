import { Component } from '@angular/core';
// import { GetKeycodes } from 'wailsjs/go/keycodes/KeycodeHelper';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  standalone: false,
})
export class AppComponent {
  // title = 'frontend';
  // keycodes: Record<number, string> = {};

  constructor() {

    // GetKeycodes().then((keycodes) => {
    //   this.keycodes = keycodes;
    // });
  }


  // GetPlaybackDeviceInfo().then((devices) => {
  //   this.devices = devices;
  // });
}
