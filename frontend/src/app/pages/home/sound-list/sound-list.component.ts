import { Component, inject, ElementRef, ChangeDetectorRef, ViewChild, AfterViewChecked } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem } from 'src/app/core/menu/menu-list/menu-list.component';
import { KeycodeService } from 'src/app/shared/keycode.service';
import { SoundHotkeysService } from 'src/app/shared/sound-hotkeys.service';

@Component({
  selector: 'sound-list',
  templateUrl: './sound-list.component.html',
  styleUrls: ['./sound-list.component.scss'],
  standalone: false,
})
export class SoundListComponent implements AfterViewChecked {
  private readonly HEADER_PADDING_NO_SCROLL = '5';
  // TODO: this is actually HEADER_PADDING_NO_SCROLL + scrollbar width, make this dynamic based on those values to always be accurate
  private readonly HEADER_PADDING_WITH_SCROLL = '18';

  keycodeService = inject(KeycodeService);
  soundHotkeysService = inject(SoundHotkeysService);

  bodyOverflow: boolean = false;
  headerPadding: string = this.HEADER_PADDING_NO_SCROLL;

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

  constructor(private router: Router, private cdr: ChangeDetectorRef) {
    this.soundHotkeysService.updateHotkeys();
  }

  @ViewChild('body') bodyElement!: ElementRef<HTMLDivElement>;
  ngAfterViewChecked() {
    // Access the native DOM element
    this.checkOverflow();
    if (this.bodyOverflow) {
      this.headerPadding = this.HEADER_PADDING_WITH_SCROLL;
    } else {
      this.headerPadding = this.HEADER_PADDING_NO_SCROLL;
    }
    this.cdr.detectChanges();
  }

  checkOverflow() {
    this.bodyOverflow = this.bodyElement.nativeElement.scrollHeight > this.bodyElement.nativeElement.clientHeight || this.bodyElement.nativeElement.scrollWidth > this.bodyElement.nativeElement.clientWidth;
  }

  fileFromPath(path: string): string {
    return path.replace(/^.*[\\\/]/, '');
  }

  editHotkey(hotkeyID: number[] | undefined) {
    if (hotkeyID == undefined) {
      return
    }

    this.router.navigate(["/edit-hotkey", hotkeyID]);
  }

  removeSound(hotkeyID: number[] | undefined) {
    if (hotkeyID == undefined) {
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
