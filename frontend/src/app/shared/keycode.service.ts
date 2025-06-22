import { Injectable } from '@angular/core';
import { GetRawcodeGroups, GetDisplayGroups, GetJSCodeGroups } from 'wailsjs/go/keycodes/KeycodeHelper';
import { keycodes } from 'wailsjs/go/models';

@Injectable({
  providedIn: 'root',
})
export class KeycodeService {
  rawcodeGroups!: Record<keycodes.KeycodeGroup, Record<number, keycodes.Keycode>>;
  displayGroups!: Record<keycodes.KeycodeGroup, Record<string, keycodes.Keycode>>;
  jsCodeGroups!: Record<keycodes.KeycodeGroup, Record<string, keycodes.Keycode>>;

  constructor() {
    GetRawcodeGroups().then((rawcodeGroups) => {
      this.rawcodeGroups = rawcodeGroups;
    });

    GetDisplayGroups().then((displayGroups) => {
      this.displayGroups = displayGroups;
    });

    GetJSCodeGroups().then((jsCodeGroups) => {
      this.jsCodeGroups = jsCodeGroups;
    });
  }

  rawcodeToString(rawcode: number): string {
    return this.rawcodeGroups[keycodes.KeycodeGroup.ALL][rawcode]?.Display;
  }

  rawcodesToString(rawcodes: number[] | undefined): string {
    return rawcodes ? rawcodes.map((rawcode) => { return this.rawcodeToString(rawcode); }).join("+") : "";
  }

  keycodesToString(keycodes: keycodes.Keycode[] | undefined): string {
    return keycodes ? keycodes.map((keycode) => { return keycode.Display; }).join("+") : "";
  }

  keycodesToRawcodes(keycodes: keycodes.Keycode[] | undefined): number[] {
    return keycodes ? keycodes.map((keycode) => { return keycode.Rawcode; }) : [];
  }
}
