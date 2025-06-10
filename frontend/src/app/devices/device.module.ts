import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { DeviceListComponent } from './device-list/device-list.component'

@NgModule({
  declarations: [
    DeviceListComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
  ],
  exports: [
    DeviceListComponent
  ],
  providers: [],
  bootstrap: [DeviceListComponent]
})
export class DeviceModule { }
