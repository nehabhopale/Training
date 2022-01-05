import{Directive, ElementRef} from '@angular/core';

@Directive({
    selector: '.appTesting'
  })
export class TestingDirective {
    constructor(private e1:ElementRef){
        console.log("instance directie",e1)
        e1.nativeElement.style.color="pink"
    }
    // constructor(private e1:ElementRef,private renderer:Renderer2){
        //console.log("instance directie",e1)
    //     renderer.setStyle(e1.nativeElement,"color","blue")
    // }
}