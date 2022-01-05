import { Directive, ElementRef } from '@angular/core';

@Directive({
  selector: '[appTest]'
})
export class TestDirective {

  constructor(private e1:ElementRef) { 
    console.log("instance of directives:",e1)
    e1.nativeElement.style.backgroundColor="purple"
  }

}
