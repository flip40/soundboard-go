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

  rawcodeToString(rawcode: number): string {
    // TODO: Test if this breaks when using a key we haven't defined in the raw codes. Maybe I can just use a smaller KeycodeGroup to test that easily.
    return this.rawcodeGroups[keycodes.KeycodeGroup.ALL][rawcode]?.Display;
    // return this.rawcodeGroups[keycodes.KeycodeGroup.ALL][rawcode]?.Display ?? "undefined"; // try this if it does break on testing
  }

  rawcodesToString(rawcodes: number[] | undefined): string {
    return rawcodes ? rawcodes.map((rawcode) => { return this.rawcodeToString(rawcode); }).join("+") : "";
  }

  // rawcodesToKeycodes(rawcodes: number[] | undefined): string {
  //   return rawcodes ? rawcodes.map((rawcode) => { return this.rawcodeToString(rawcode); }).join("+") : "";
  // }

  keycodesToString(keycodes: keycodes.Keycode[] | undefined): string {
    return keycodes ? keycodes.map((keycode) => { return keycode.Display; }).join("+") : "";
  }

  keycodesToRawcodes(keycodes: keycodes.Keycode[] | undefined): number[] {
    return keycodes ? keycodes.map((keycode) => { return keycode.Rawcode; }) : [];
  }
}
