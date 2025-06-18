import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { HomeModule } from './home/home.module';
// import { DeviceModule } from './devices/device.module'
// import { SoundModule } from './sounds/sound.module'
import { AppHeaderComponent } from './app-header/app-header.component'

import { provideRouter, RouterModule } from '@angular/router';
import routeConfig from './routes';
import { MenuModule } from "./menu/menu.module";

@NgModule({
  declarations: [
    AppComponent,
    AppHeaderComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule,
    HomeModule,
    MenuModule
  ],
  providers: [provideRouter(routeConfig)],
  bootstrap: [AppComponent]
})
export class AppModule { }
