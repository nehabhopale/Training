import { Directive , ElementRef, HostListener } from '@angular/core';
import { NgControl } from '@angular/forms';

@Directive({
  selector: '[appChecker]'
})
export class CheckerDirective {

  constructor(private el: ElementRef,private control: NgControl) { 
    // this.control.viewToModelUpdate((target.value === '') ? null : target.value);
    if (this.el.nativeElement.value===""){
      this.el.nativeElement.value=null
      
    }
  }


@HostListener('input', ['$event.target']) //to listen things happening on dom element
  onEvent(target: HTMLInputElement){
    this.control.viewToModelUpdate((target.value === '') ? null : target.value);
    // console.log(target.value)
    }
}