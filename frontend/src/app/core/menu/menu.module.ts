import { NgModule } from '@angular/core';
import { MenuListComponent } from './menu-list/menu-list.component';
import { MenuGroupComponent } from './menu-group/menu-group.component';

@NgModule({
  declarations: [
    MenuGroupComponent,
    MenuListComponent,
  ],
  imports: [],
  exports: [
    MenuGroupComponent,
    MenuListComponent,
  ],
  providers: [],
})
export class MenuModule { }
