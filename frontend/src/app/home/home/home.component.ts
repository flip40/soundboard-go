import { Component, inject } from '@angular/core';
import { AddSounds } from 'wailsjs/go/main/App'
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys/sound-hotkeys.service';

@Component({
    selector: 'home',
    templateUrl: './home.component.html',
    styleUrls: ['./home.component.scss'],
    standalone: false,
})
export class HomeComponent {
  soundhotkeysService = inject(SoundHotkeysService);

  constructor() {

  }

  addSounds() {
    AddSounds().then(() => {
      this.soundhotkeysService.updateHotkeys();
    });
  }

  setStopHotkey() {

  }
}
