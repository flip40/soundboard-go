import { Component, inject, signal, HostListener } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { KeycodeService } from 'src/app/shared/keycodes/keycode.service';
import { keycodes } from 'wailsjs/go/models';
// import { AddSounds } from 'wailsjs/go/main/App'
// import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys/sound-hotkeys.service';

@Component({
  selector: 'hotkey-editor',
  templateUrl: './hotkey-editor.component.html',
  styleUrls: ['./hotkey-editor.component.scss'],
  standalone: false,
})
export class HotkeyEditorComponent {
  keycodeService = inject(KeycodeService);

  hotkeyID = signal('');
  private activatedRoute = inject(ActivatedRoute);

  constructor(private router: Router) {
    this.activatedRoute.params.subscribe((params) => {
      this.hotkeyID.set(params['id']);
    });
  }

  @HostListener('document:keydown', ['$event'])
  onKeyDown(event: KeyboardEvent) {
    event.preventDefault();

    // TODO: finish handling for getting user hotkey input
    console.log('event.key:', event.key, 'event.code:', event.code);
    // if (event.key === 'Enter') {
    if (this.keycodeService.jsCodeGroups[keycodes.KeycodeGroup.MODIFIERS][event.code] != undefined) {
      console.log('Modifier pressed');
    }
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
