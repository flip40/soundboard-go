import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HomeComponent } from './home/home.component';
import { DeviceModule } from '../devices/device.module'
import { SoundModule } from '../sounds/sound.module'

// import { provideRouter } from '@angular/router';
// import routeConfig from './routes';

@NgModule({
  declarations: [
    HomeComponent
  ],
  imports: [
    BrowserModule,
    DeviceModule,
    SoundModule,
  ],
  exports: [
    HomeComponent
  ],
  providers: [],
  bootstrap: [HomeComponent]
})
export class HomeModule { }
