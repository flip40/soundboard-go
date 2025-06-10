import { Injectable } from '@angular/core';
import { GetKeycodes } from 'wailsjs/go/keycodes/KeycodeHelper';

@Injectable({
  providedIn: 'root',
})
export class KeycodeService {
  keycodes: Record<number, string> = {};

  constructor() {
    GetKeycodes().then((keycodes) => {
      this.keycodes = keycodes;
    });
  }

  getKeycodes(): Record<number, string> {
    return this.keycodes;
  }

  rawcodeToString(rawcode: number): string {
    return this.keycodes[rawcode];
  }
}
