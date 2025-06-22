import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';

import { AppHeaderModule } from './core/app-header/app-header.module';
import { HomeModule } from './pages/home/home.module';
import { MenuModule } from "./core/menu/menu.module";

import { provideRouter, RouterModule } from '@angular/router';
import routeConfig from './routes';

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule,
    AppHeaderModule,
    HomeModule,
    MenuModule,
  ],
  providers: [provideRouter(routeConfig)],
  bootstrap: [AppComponent]
})
export class AppModule { }
