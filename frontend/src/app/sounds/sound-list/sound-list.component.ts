import { Component, Input, inject } from '@angular/core';
import { GetSoundHotkeys, ShowDialog } from 'wailsjs/go/main/App';
import { soundhotkey } from 'wailsjs/go/models';
import { KeycodeService } from 'src/app/shared/keycodes/keycode.service';
// import { HotkeyToString } from 'wailsjs/go/soundhotkey/HotkeyHelper';

// TODO: This almost certainly should be passed in by App initially and updated in a different way...
// import { GetPlaybackDeviceInfo } from "../../../../wailsjs/go/main/App"
// import { audiodevice } from 'wailsjs/go/models';

@Component({
    selector: 'sound-list',
    templateUrl: './sound-list.component.html',
    styleUrls: ['./sound-list.component.scss'],
    standalone: false,
})
export class SoundListComponent {
  keycodeService = inject(KeycodeService);
  // @Input() keycodes: Record<number, string> = {};
  // @Input() extDevices = [];

  soundHotkeys: soundhotkey.SoundHotkey[] = [];
  // selectedDevice: number[] = [];

  constructor() {
    
    // TODO: This almost certainly should be passed in by App initially and updated in a different way...
    GetSoundHotkeys().then((soundHotkeys) => {
      this.soundHotkeys = soundHotkeys;
    });
  }

  fileFromPath(path: string): string {
    return path.replace(/^.*[\\\/]/, '');
  }

  hotkeyToString(hotkey: number[]): string {
    // TODO: hotkey is probably going to need to be an array unless we do the conversion here... Maybe we can handle a lot of this by
    //       sending hotkey updates directly to Go, though
    // return hotkey ? hotkey.map((keyCode: KeyCode) => { return keyCode.electronCode; }).join("+") : ""
    // let str = await HotkeyToString(hotkey).then((result) => {
    //   return result;
    // })

    let keys: string[] = [];
    hotkey.forEach((rawcode: number) => {
      keys.push(this.keycodeService.rawcodeToString(rawcode))
    });

    // return HotkeyToString(hotkey)
    return keys.join("+");
  }

  editHotkey(hotkey: number[]) {

    // TODO: DEBUG
    ShowDialog("double clicked!");
  }

  openContextMenu(event: Event, hotkey: number[]) {
    
    // TODO: DEBUG
    ShowDialog("right clicked!");
  }
}
