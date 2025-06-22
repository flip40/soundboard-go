import { NgModule } from '@angular/core';
import { AppHeaderComponent } from './app-header.component';
import { MenuModule } from '../menu/menu.module';

@NgModule({
  declarations: [
    AppHeaderComponent,
  ],
  imports: [
    MenuModule,
  ],
  exports: [
    AppHeaderComponent
  ],
  providers: [],
})
export class AppHeaderModule { }
