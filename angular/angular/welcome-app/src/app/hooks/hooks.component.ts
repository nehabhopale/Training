import { Component, OnChanges, OnInit, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-hooks',
  templateUrl: './hooks.component.html',
  styleUrls: ['./hooks.component.css']
})
export class HooksComponent implements OnInit,OnChanges{
  itemsOfHook:any=20
  flag:boolean=false
  constructor() { 
    console.log("inside hooks constructor")
  }

  ngOnInit(): void {
    console.log("inside hooks init")
  }
  ngOnChanges(changes:SimpleChanges):void{
    console.log("in onchange of parent",changes)
  }
  ngOnDestroy(): void {
    console.log("Parent destroyed")
  }
  ngDoCheck(): void {
    console.log("calling docheck of parent")
    // parent called after ngOnchanges and first time after onint
}
  hookHandler(e:any){
    console.log("Parent hadnler ",e )
  }

}
