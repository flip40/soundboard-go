import { Injectable } from '@angular/core';
import { GetRawcodeGroups, GetDisplayGroups, GetJSCodeGroups } from 'wailsjs/go/keycodes/KeycodeHelper';
import { keycodes } from 'wailsjs/go/models';

@Injectable({
  providedIn: 'root',
})
export class KeycodeService {
  // keycodes: Record<number, string> = {};
  rawcodeGroups!: Record<keycodes.KeycodeGroup, Record<number, keycodes.Keycode>>;
  displayGroups!: Record<keycodes.KeycodeGroup, Record<string, keycodes.Keycode>>;
  jsCodeGroups!: Record<keycodes.KeycodeGroup, Record<string, keycodes.Keycode>>;

  constructor() {
    // GetKeycodes().then((keycodes) => {
    //   this.keycodes = keycodes;
    // });

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

  // getKeycodes(): Record<number, string> {
  //   return this.keycodes;
  // }

  // getKeycodeGroups(): Record<keycodes.KeycodeGroup, Record<number, string>> {
  //   return this.keycodeGroups;
  // }

  // getKeycodeStringGroups(): Record<keycodes.KeycodeGroup, Record<string, number>> {
  //   return this.keycodeStringGroups;
  // }

  rawcodeToString(rawcode: number): string {
    return this.rawcodeGroups[keycodes.KeycodeGroup.ALL][rawcode].Display;
  }
}
