import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys.service';
import { AddSounds } from 'wailsjs/go/main/App'

@Component({
  selector: 'home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  standalone: false,
})
export class HomeComponent {
  soundhotkeysService = inject(SoundHotkeysService);

  constructor(private router: Router) {

  }

  addSounds() {
    AddSounds().then(() => {
      this.soundhotkeysService.updateHotkeys();
    });
  }

  setStopHotkey() {
    this.router.navigate(["/edit-hotkey", "stop"], {
      queryParams: {
        isStop: true,
      }
    });
  }
}
