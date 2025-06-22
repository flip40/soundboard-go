import { Component, HostListener, inject } from '@angular/core';
import { KeycodeService } from 'src/app/shared/keycode.service';
import { keycodes } from 'wailsjs/go/models';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  standalone: false,
})
export class AppComponent {
  keycodeService: KeycodeService = inject(KeycodeService);

  constructor() { }

  @HostListener('document:keydown', ['$event'])
  onKeyDown(event: KeyboardEvent) {
    // Ignore Ctrl+A
    if (event.ctrlKey && event.code == this.keycodeService.displayGroups[keycodes.KeycodeGroup.ALL]["A"].JSCode) {
      event.preventDefault();
    }
  }
}
