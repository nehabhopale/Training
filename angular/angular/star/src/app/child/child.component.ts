import { Component, OnInit,Output,Input,EventEmitter } from '@angular/core';

@Component({
  selector: 'app-child',
  templateUrl: './child.component.html',
  styleUrls: ['./child.component.css']
})
export class ChildComponent implements OnInit {
  // starRating = 0; 
  // @Output() childRating:EventEmitter<number>=new EventEmitter()
  @Input() rating!: number
  visible!:string
  stars: number[] = [1,2, 3, 4, 5];
  constructor() { }

  ngOnInit(): void {
  }
  getWidth():string{
    return this.rating*41.8+'px'
  }
  getOuterDivWidth():string{
    
    let w=this.rating*80
    this.visible="hidden"
    console.log("outer",w)
    return w +'px'

  }
  getInnerDivWidth():string{

    return '400px'
  }
  countStar(star:any) {
    console.log("in count")

    this.rating = star;
    console.log('Value of star', this.rating);
  }
  // getInnerDivWidth():string{

  //   return '150px'
  // }
 

}
