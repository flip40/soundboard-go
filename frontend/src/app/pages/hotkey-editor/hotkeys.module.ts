import { NgModule } from '@angular/core';

import { HotkeyEditorComponent } from './hotkey-editor.component';
// import { DeviceModule } from '../devices/device.module'
// import { SoundModule } from '../sounds/sound.module'

// import { provideRouter } from '@angular/router';
// import routeConfig from './routes';

@NgModule({
  declarations: [
    HotkeyEditorComponent,
  ],
  imports: [],
  exports: [
    HotkeyEditorComponent
  ],
  providers: [],
})
export class HotkeysModule { }
