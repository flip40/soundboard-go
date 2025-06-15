import { Component, Input, inject, signal, WritableSignal, HostListener, ElementRef, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { ShowDialog } from 'wailsjs/go/main/App';
import { KeycodeService } from 'src/app/shared/keycode.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys.service';
import { MenuItem, MenuListComponent } from 'src/app/menu/menu-list/menu-list.component';
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
  contextMenuX: number = 0;
  contextMenuY: number = 0;
  contextMenuActive: boolean = false;
  contextMenuHotkey: number[] | undefined;
  contextMenuItems: MenuItem[] = [
    {
      text: "Edit Hotkey...",
      onClick: () => this.editHotkey(this.contextMenuHotkey),
    },
    {
      text: "Remove Sound",
      onClick: () => this.removeSound(this.contextMenuHotkey),
    },
  ]

  // soundHotkeys: soundhotkey.SoundHotkey[] = [];
  // selectedDevice: number[] = [];

  constructor(private router: Router) {
    this.soundHotkeysService.updateHotkeys();
  }

  fileFromPath(path: string): string {
    return path.replace(/^.*[\\\/]/, '');
  }

  editHotkey(hotkeyID: number[] | undefined) {
    if (hotkeyID == undefined) {
      // TODO: error?
      return
    }
    this.router.navigate(["/edit-hotkey", hotkeyID]);
  }

  removeSound(hotkeyID: number[] | undefined) {
    if (hotkeyID == undefined) {
      // TODO: error?
      return
    }
    this.soundHotkeysService.removeSound(hotkeyID);
    this.contextMenuActive = false;
  }

  openContextMenu(event: MouseEvent, hotkeyID: number[]) {
    this.contextMenuX = event.clientX;
    this.contextMenuY = event.clientY;
    this.contextMenuActive = true;
    this.contextMenuHotkey = hotkeyID;
  }

  closeContextMenu() {
    this.contextMenuActive = false;
    this.contextMenuHotkey = undefined;
  }
}
