import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MenuListComponent } from './menu-list/menu-list.component';
import { MenuGroupComponent } from './menu-group/menu-group.component';

@NgModule({
  declarations: [
    MenuGroupComponent,
    MenuListComponent,
  ],
  imports: [
    BrowserModule,
  ],
  exports: [
    MenuGroupComponent,
    MenuListComponent,
  ],
  providers: [],
  bootstrap: [MenuListComponent, MenuGroupComponent]
})
export class MenuModule { }
