import { Routes } from '@angular/router';

import { HomeComponent } from './pages/home/home.component';
import { HotkeyEditorComponent } from './pages/hotkey-editor/hotkey-editor.component';

const routeConfig: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: 'Home',
  },
  {
    path: 'edit-hotkey/:id',
    component: HotkeyEditorComponent,
    title: 'Hotkey Editor',
  },
];
export default routeConfig;