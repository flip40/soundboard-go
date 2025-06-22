import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem } from '../menu/menu-list/menu-list.component'
import { MenuGroup } from '../menu/menu-group/menu-group.component'
import { AudioDeviceService } from 'src/app/shared/audio-device.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys.service';
import { WindowService } from 'src/app/shared/window.service';
import { ResetSoundboard, LoadSoundboard, SaveSoundboard } from 'wailsjs/go/main/App';
import { Quit, WindowMinimise } from 'wailsjs/runtime/runtime';

@Component({
  selector: 'app-header',
  templateUrl: './app-header.component.html',
  styleUrls: ['./app-header.component.scss'],
  standalone: false,
})
export class AppHeaderComponent {
  audioDeviceService = inject(AudioDeviceService);
  soundHotkeysService = inject(SoundHotkeysService);
  windowService = inject(WindowService);

  menuActive: boolean = false;
  activeGroup: number = 0;

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
    },
    // {
    //   title: "Edit",
    //   menuList: [
    //     {
    //       text: "blerh...",
    //       onClick: () => this.newSoundboard(),
    //     },
    //   ],
    // },
  ]

  constructor(private router: Router) { }

  onMouseEnter(index: number) {
    this.activeGroup = index;
  }

  // TODO: handle menu is active and hover state changes to another menu group
  toggleMenu() {
    this.menuActive = !this.menuActive;
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
      this.router.navigate([""]);
    });
  }

  openSoundboard() {
    LoadSoundboard().then(() => {
      this.audioDeviceService.updateAudioDevices();
      this.soundHotkeysService.updateHotkeys();
      this.deactivteMenu();
      this.router.navigate([""]);
    });
  }

  saveSoundboard() {
    SaveSoundboard().then(() => {
      this.deactivteMenu();
    });
  }

  minimize() {
    // TODO: need to remove hover state when un-minimizing
    WindowMinimise();
  }

  closeApp() {
    // TODO: may want to send this to the go app instead to handle "save before close"
    Quit();
  }
}
