import { Component, HostListener, inject } from '@angular/core';
import { KeycodeService } from 'src/app/shared/keycode.service';
import { keycodes } from 'wailsjs/go/models';
// import { GetKeycodes } from 'wailsjs/go/keycodes/KeycodeHelper';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  standalone: false,
})
export class AppComponent {
  keycodeService: KeycodeService = inject(KeycodeService);
  // title = 'frontend';
  // keycodes: Record<number, string> = {};

  constructor() {
    // GetKeycodes().then((keycodes) => {
    //   this.keycodes = keycodes;
    // });
  }

  @HostListener('document:keydown', ['$event'])
  onKeyDown(event: KeyboardEvent) {
    // Ignore Ctrl+A
    if (event.ctrlKey && event.code == this.keycodeService.displayGroups[keycodes.KeycodeGroup.ALL]["A"].JSCode) {
      event.preventDefault();
    }
  }


  // GetPlaybackDeviceInfo().then((devices) => {
  //   this.devices = devices;
  // });
}
