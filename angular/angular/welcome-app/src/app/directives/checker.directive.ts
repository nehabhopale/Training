import { Directive, ElementRef, HostListener } from '@angular/core';
// import { NgControl } from '@angular/forms';
// import {NgControl} from "@angular/common";
@Directive({
  selector: '[appChecker]'
})
export class CheckerDirective {

  constructor(private el: ElementRef) { 

  }
  @HostListener('change', )
  onChange(){
    console.log("changed")
    alert(this.el.nativeElement.value)
    if (this.el.nativeElement.value===""){
      // console.log("value null")
      this.el.nativeElement.value=undefined  
      
    }
    // alert(this.el.nativeElement.value)
    console.log(this.el.nativeElement.value)
    
  }
  // @HostListener('input', ['$event.target'])
  // onEvent(target: HTMLInputElement){
  //   this.control.viewToModelUpdate((target.value === '') ? null : target.value);
  //   console.log(target.value)
    
  // }

}
