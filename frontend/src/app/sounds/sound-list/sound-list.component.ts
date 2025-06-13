import { Component, Input, inject } from '@angular/core';
import { Router } from '@angular/router';
import { ShowDialog } from 'wailsjs/go/main/App';
import { KeycodeService } from 'src/app/shared/keycodes/keycode.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys/sound-hotkeys.service';
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
  soundHotkeysService = inject(SoundHotkeysService);
  // @Input() keycodes: Record<number, string> = {};
  // @Input() extDevices = [];

  // soundHotkeys: soundhotkey.SoundHotkey[] = [];
  // selectedDevice: number[] = [];

  constructor(private router: Router) {
    this.soundHotkeysService.updateHotkeys();
  }

  fileFromPath(path: string): string {
    return path.replace(/^.*[\\\/]/, '');
  }

  editHotkey(hotkeyID: number[]) {

    // TODO: DEBUG
    // ShowDialog("double clicked! id: " + hotkeyID);
    this.router.navigate(["/edit-hotkey", hotkeyID]);
  }

  openContextMenu(event: Event, hotkeyID: number[]) {

    // TODO: DEBUG
    ShowDialog("right clicked! id: " + hotkeyID);
  }
}
