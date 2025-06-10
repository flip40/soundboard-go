import { Component, inject, signal } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
// import { AddSounds } from 'wailsjs/go/main/App'
// import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys/sound-hotkeys.service';

@Component({
  selector: 'hotkey-editor',
  templateUrl: './hotkey-editor.component.html',
  styleUrls: ['./hotkey-editor.component.scss'],
  standalone: false,
})
export class HotkeyEditorComponent {
  hotkeyID = signal('');
  private activatedRoute = inject(ActivatedRoute);

  constructor(private router: Router) {
    this.activatedRoute.params.subscribe((params) => {
      this.hotkeyID.set(params['id']);
    });
  }

  toHome() {
    this.router.navigate([""]);
  }

  //   addSounds() {
  //     AddSounds().then(() => {
  //       this.soundhotkeysService.updateHotkeys();
  //     });
  //   }

  //   setStopHotkey() {

  //   }
}
