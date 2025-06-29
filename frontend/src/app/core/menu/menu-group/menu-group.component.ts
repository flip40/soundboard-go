import { Component, Input, Output, EventEmitter, ElementRef } from '@angular/core';
import { MenuItem } from '../menu-list/menu-list.component';

export interface MenuGroup {
  title: string;
  menuList: MenuItem[];
}

@Component({
  selector: 'menu-group',
  templateUrl: './menu-group.component.html',
  styleUrls: ['./menu-group.component.scss'],
  standalone: false,
})
export class MenuGroupComponent {
  @Input() title: string = "";
  @Input() menuList: MenuItem[] = [];
  @Input() isActive: boolean = false;
  @Output() notClicked = new EventEmitter<void>();
  menuX: number = 0;
  menuY: number = 0;

  constructor(private elementRef: ElementRef) { }

  menuNotClicked(event: MouseEvent) {
    // Don't handle if other menu group item was clicked
    if (!this.isActive || this.elementRef.nativeElement.contains(event.target)) {
      return;
    }

    this.notClicked.emit();
  }

}
