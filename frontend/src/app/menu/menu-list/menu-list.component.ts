import { Component, Input, Output, EventEmitter, HostBinding, ElementRef, HostListener } from '@angular/core';

export interface MenuItem {
  text: string;
  children?: MenuItem[];
  onClick(): void;
}

@Component({
  selector: 'menu-list',
  templateUrl: './menu-list.component.html',
  styleUrls: ['./menu-list.component.scss'],
  standalone: false,
})
export class MenuListComponent {
  @Input() menuItems: MenuItem[] = [];
  @Input() isActive: boolean = false;
  @HostBinding('style.left.px') @Input() leftPos: number = 0;
  @HostBinding('style.top.px') @Input() topPos: number = 0;
  @HostBinding('class.hide') get shouldHide(): boolean { return !this.isActive; };
  @Output() notClicked = new EventEmitter<void>();

  constructor(private elementRef: ElementRef) {

  }

  // Handle any click that isn't in the menu to close the menu
  @HostListener('window:mousedown', ['$event'])
  onMousedown(event: Event): void {
    if (this.elementRef.nativeElement.contains(event.target)) {
      return;
    }

    this.notClicked.emit();
  }

}
