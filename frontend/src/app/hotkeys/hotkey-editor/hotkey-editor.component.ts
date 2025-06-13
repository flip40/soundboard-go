import { Component, inject, signal, HostListener, WritableSignal } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { KeycodeService } from 'src/app/shared/keycodes/keycode.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys/sound-hotkeys.service';
import { keycodes, soundhotkey } from 'wailsjs/go/models';
import { debounce, BetterSet } from 'src/utils/helpers'
import { SetHotkey, ClearHotkey } from 'wailsjs/go/main/App';
// import { AddSounds } from 'wailsjs/go/main/App'
// import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys/sound-hotkeys.service';

@Component({
  selector: 'hotkey-editor',
  templateUrl: './hotkey-editor.component.html',
  styleUrls: ['./hotkey-editor.component.scss'],
  standalone: false,
})
export class HotkeyEditorComponent {
  keycodeService: KeycodeService = inject(KeycodeService);
  soundHotkeysService = inject(SoundHotkeysService);

  soundHotkeyID: WritableSignal<number[]> = signal([]);
  soundHotkey: soundhotkey.SoundHotkey | undefined;
  oldHotkey: number[] | undefined = [];
  newHotkey: keycodes.Keycode[] = [];

  private activatedRoute: ActivatedRoute = inject(ActivatedRoute);
  private capturing: boolean = false;
  private keysPressed: BetterSet<string> = new BetterSet<string>();

  constructor(private router: Router) {
    this.activatedRoute.params.subscribe((params) => {
      this.soundHotkeyID.set(params['id']);
      this.soundHotkey = this.soundHotkeysService.getHotkeyByID(this.soundHotkeyID());
      if (this.soundHotkey == undefined) {
        // TODO: ERROR, this should never happen
      }
      this.oldHotkey = this.soundHotkey?.Hotkey;
    });
  }

  @HostListener('document:keydown', ['$event'])
  onKeyDown(event: KeyboardEvent) {
    event.preventDefault();

    if (!this.capturing) {
      this.capturing = true;
      this.getCaptured();
    }

    this.keysPressed.add(event.code);
  }


  private getCaptured = debounce(() => {
    this.capturing = false;

    let modifiers: keycodes.Keycode[] = new Array<keycodes.Keycode>();
    let selectedKey: keycodes.Keycode | undefined = undefined;

    this.keysPressed.forEach((key: string) => {
      if (this.keycodeService.jsCodeGroups[keycodes.KeycodeGroup.MODIFIERS][key] != undefined) {
        modifiers.push(this.keycodeService.jsCodeGroups[keycodes.KeycodeGroup.MODIFIERS][key]);
      } else if (this.keycodeService.jsCodeGroups[keycodes.KeycodeGroup.ALL][key]) {
        selectedKey = this.keycodeService.jsCodeGroups[keycodes.KeycodeGroup.ALL][key];
      } else {
        console.log("key not defined in keycodes:", key);
      }
    });

    if (selectedKey === undefined) {
      // error, don't set a hotkey
      this.keysPressed = new BetterSet<string>();
      return;
    }

    // this.setState({ hotkey: [...modifiers.sort(), selectedKey].join("+") });
    this.newHotkey = [...modifiers.sort((a: keycodes.Keycode, b: keycodes.Keycode) => a.Display.localeCompare(b.Display)), selectedKey];
    this.keysPressed = new BetterSet<string>();
  }, 100);

  setHotkey() {
    console.log("this.soundHotkeyID:", this.soundHotkeyID());
    SetHotkey(String(this.soundHotkeyID()), this.keycodeService.keycodesToRawcodes(this.newHotkey)).then(() => {
      // TODO: Debug
      console.log("set hotkey");
      this.goHome();
    });
  }

  clearHotkey() {
    ClearHotkey(String(this.soundHotkeyID())).then(() => {
      // TODO: Debug
      console.log("cleared hotkey");
      this.goHome();
    });
  }

  goHome() {
    this.router.navigate([""]);
  }
}
