import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { HomeModule } from './home/home.module';
// import { DeviceModule } from './devices/device.module'
// import { SoundModule } from './sounds/sound.module'

import { provideRouter, RouterModule } from '@angular/router';
import routeConfig from './routes';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    RouterModule,
    HomeModule,
    // DeviceModule,
    // SoundModule,
  ],
  providers: [provideRouter(routeConfig)],
  bootstrap: [AppComponent]
})
export class AppModule { }
