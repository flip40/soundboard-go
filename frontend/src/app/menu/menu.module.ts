import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MenuListComponent } from './menu-list/menu-list.component';

@NgModule({
  declarations: [
    MenuListComponent,
  ],
  imports: [
    BrowserModule,
  ],
  exports: [
    MenuListComponent
  ],
  providers: [],
  bootstrap: [MenuListComponent]
})
export class MenuModule { }
