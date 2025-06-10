import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { SoundListComponent } from './sound-list/sound-list.component'

@NgModule({
  declarations: [
    SoundListComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
  ],
  exports: [
    SoundListComponent
  ],
  providers: [],
  bootstrap: [SoundListComponent]
})
export class SoundModule { }
