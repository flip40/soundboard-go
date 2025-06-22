import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MenuModule } from 'src/app/core/menu/menu.module';

import { HomeComponent } from './home.component';
import { DeviceListComponent } from './device-list/device-list.component';
import { SoundListComponent } from './sound-list/sound-list.component';


@NgModule({
  declarations: [
    HomeComponent,
    DeviceListComponent,
    SoundListComponent
  ],
  imports: [
    FormsModule,
    MenuModule,
  ],
  exports: [
    HomeComponent
  ],
  providers: [],
})
export class HomeModule { }
