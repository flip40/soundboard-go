import { Component, inject, signal, HostListener, WritableSignal } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { KeycodeService } from 'src/app/shared/keycode.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys.service';
import { debounce, BetterSet } from 'src/utils/helpers'
import { SetHotkey, SetStopHotkey, ClearHotkey, ClearStopHotkey, ErrorDialog } from 'wailsjs/go/app/App';
import { keycodes, soundhotkey } from 'wailsjs/go/models';

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
  isStop: boolean = false;

  private activatedRoute: ActivatedRoute = inject(ActivatedRoute);
  private capturing: boolean = false;
  private keysPressed: BetterSet<string> = new BetterSet<string>();

  constructor(private router: Router) {
    const snapshot = this.activatedRoute.snapshot;
    this.isStop = snapshot.queryParamMap.has('isStop') ? Boolean(snapshot.queryParamMap.get('isStop')) : false;

    if (this.isStop) {
      this.oldHotkey = this.soundHotkeysService.getStopHotkey();
    } else {
      this.soundHotkeyID.set(snapshot.params['id']);
      this.soundHotkey = this.soundHotkeysService.getHotkeyByID(this.soundHotkeyID());
      if (this.soundHotkey == undefined) {
        ErrorDialog("Invalid Sound Hotkey!");
        this.router.navigate([""]);
      }
      this.oldHotkey = this.soundHotkey?.hotkey;
    }
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

    this.newHotkey = [...modifiers.sort((a: keycodes.Keycode, b: keycodes.Keycode) => a.Display.localeCompare(b.Display)), selectedKey];
    this.keysPressed = new BetterSet<string>();
  }, 100);

  // TODO: Consider moving SetHotkey calls to hotkey service
  setHotkey() {
    if (this.isStop) {
      SetStopHotkey(this.keycodeService.keycodesToRawcodes(this.newHotkey)).then(() => {
        this.goHome();
      });
    } else {
      SetHotkey(String(this.soundHotkeyID()), this.keycodeService.keycodesToRawcodes(this.newHotkey)).then(() => {
        this.goHome();
      });
    }
  }

  // TODO: Consider moving ClearHotkey calls to hotkey service
  clearHotkey() {
    if (this.isStop) {
      ClearStopHotkey().then(() => {
        this.goHome();
      });
    } else {
      ClearHotkey(String(this.soundHotkeyID())).then(() => {
        this.goHome();
      });

    }
  }

  goHome() {
    this.router.navigate([""]);
  }
}
