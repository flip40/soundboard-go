import { Injectable, signal } from '@angular/core';
import { WindowIsMaximised, WindowMaximise, WindowUnmaximise } from 'wailsjs/runtime/runtime';

@Injectable({
  providedIn: 'root',
})
export class WindowService {
  isMaximized = signal<boolean>(false);

  constructor() {
    this.updateWindowState();

    // TODO: might want to debounce this if weird behavior starts, it triggers a lot.
    window.addEventListener('resize', this.updateWindowState.bind(this));
  }

  updateWindowState() {
    WindowIsMaximised().then((isMaximized) => {
      this.isMaximized.set(isMaximized);
    });
  }

  toggleMaximize() {
    if (this.isMaximized()) {
      WindowUnmaximise();
      this.isMaximized.set(false);
    } else {
      WindowMaximise();
      this.isMaximized.set(true);
    }
  }
}
