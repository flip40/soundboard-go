import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { MenuModule } from '../menu/menu.module'

import { SoundListComponent } from './sound-list/sound-list.component'

@NgModule({
  declarations: [
    SoundListComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    MenuModule,
  ],
  exports: [
    SoundListComponent
  ],
  providers: [],
  bootstrap: [SoundListComponent]
})
export class SoundModule { }
