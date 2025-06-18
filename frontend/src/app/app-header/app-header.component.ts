import { Component, inject, Output, EventEmitter, HostBinding, ElementRef, HostListener } from '@angular/core';
import { MenuItem } from '../menu/menu-list/menu-list.component'
import { MenuGroup } from '../menu/menu-group/menu-group.component'
import { ResetSoundboard, LoadSoundboard, SaveSoundboard } from 'wailsjs/go/main/App';
import { AudioDeviceService } from 'src/app/shared/audio-device.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys.service';

@Component({
  selector: 'app-header',
  templateUrl: './app-header.component.html',
  // styleUrls: ['./app-header.component.scss'],
  standalone: false,
})
export class AppHeaderComponent {
  audioDeviceService = inject(AudioDeviceService);
  soundHotkeysService = inject(SoundHotkeysService);
  // @Input() menuItems: MenuItem[] = [];
  // @Input() isActive: boolean = false;
  // @HostBinding('style.left.px') @Input() leftPos: number = 0;
  // @HostBinding('style.top.px') @Input() topPos: number = 0;
  // @HostBinding('class.hide') get shouldHide(): boolean { return !this.isActive; };
  // @Output() notClicked = new EventEmitter<void>();
  fileMenu: MenuItem[] = [
    {
      text: "New Soundboard...",
      onClick: () => this.newSoundboard(),
    },
    {
      text: "Open Soundboard...",
      onClick: () => this.openSoundboard(),
    },
    // TODO: Split Save and SaveAs
    {
      text: "Save...",
      onClick: () => this.saveSoundboard(),
    },
    {
      text: "Save As...",
      onClick: () => this.saveSoundboard(),
    },
  ]

  menuGroups: MenuGroup[] = [
    {
      title: "File",
      menuList: this.fileMenu,
    }
  ]

  menuActive: boolean = false;
  activeGroup: number = 0;

  constructor() {

  }

  onMouseEnter(index: number) {
    this.activeGroup = index;
  }

  activateMenu() {
    this.menuActive = true;
  }

  deactivteMenu() {
    this.activeGroup = 0;
    this.menuActive = false;
  }

  isGroupActive(index: number) {
    return this.menuActive && this.activeGroup == index;
  }

  newSoundboard() {
    ResetSoundboard().then(() => {
      this.audioDeviceService.updateAudioDevices();
      this.soundHotkeysService.updateHotkeys();
      this.deactivteMenu();
    });
  }

  openSoundboard() {
    LoadSoundboard().then(() => {
      this.audioDeviceService.updateAudioDevices();
      this.soundHotkeysService.updateHotkeys();
      this.deactivteMenu();
    });
  }

  saveSoundboard() {
    SaveSoundboard().then(() => {
      this.deactivteMenu();
    });
  }
}
