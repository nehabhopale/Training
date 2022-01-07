import { Component, Output,EventEmitter,OnInit,Input, OnChanges, SimpleChange, SimpleChanges, OnDestroy, ContentChild,ViewChild} from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent implements OnInit,OnChanges,OnDestroy {
  @Input() item:any=0
  @Output() clickChild:EventEmitter<any>=new EventEmitter()
 
  constructor() {
    console.log("inside child constructor")
    console.log("parent template",this.varHookHello)
    console.log("child template",this.childHello)

   }
   @ContentChild('hookHello',{static:true})varHookHello:any
   @ViewChild('childHello',{static:true})childHello:any
 
  ngOnInit(): void {
    console.log("inside child init")
    console.log("parent template",this.varHookHello)
    console.log("child template",this.childHello)
  }
  ngOnChanges(changes:SimpleChanges):void{
    console.log("in onchange of child",changes)
  }
  ngOnDestroy(){
    console.log("on destroy in child ")
  }
  ngAfterContentInit():void{
    console.log("inside after content init")
    console.log("value of vookHello in after content",this.varHookHello)
    console.log("value of child hello in after content",this.childHello)
  }
  ngDoCheck(): void {
    console.log("do check on child")
  // Called immediately after ngOnChanges() on every change detection 
  }
  ngAfterViewInit():void{
    console.log("inside after content view")
    console.log("value of vookHello in after view",this.varHookHello)
    console.log("value of child hello in after view",this.childHello)
    // child is not undefined after using view int even if static not used
  }
  handleClick():void{
    // console.log("p clicked --------")
    this.clickChild.emit("event emit")
  }
  

}
