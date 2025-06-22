import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppHeaderComponent } from './app-header.component';
import { MenuModule } from '../../menu/menu.module';

@NgModule({
  declarations: [
    AppHeaderComponent,
  ],
  imports: [
    BrowserModule,
    MenuModule,
  ],
  exports: [
    AppHeaderComponent
  ],
  providers: [],
  bootstrap: [AppHeaderComponent]
})
export class AppHeaderModule { }
