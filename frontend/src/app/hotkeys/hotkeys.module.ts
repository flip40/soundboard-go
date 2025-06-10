import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HotkeyEditorComponent } from './hotkey-editor/hotkey-editor.component';
// import { DeviceModule } from '../devices/device.module'
// import { SoundModule } from '../sounds/sound.module'

// import { provideRouter } from '@angular/router';
// import routeConfig from './routes';

@NgModule({
  declarations: [
    HotkeyEditorComponent,
  ],
  imports: [
    BrowserModule,
    // DeviceModule,
    // SoundModule,
  ],
  exports: [
    HotkeyEditorComponent
  ],
  providers: [],
  bootstrap: [HotkeyEditorComponent]
})
export class HotkeysModule { }
