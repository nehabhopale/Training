import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'star-rating';
  stars: number[] = [1, 2, 3, 4, 5];
  selectedValue!: number;
  starRating = 0; 
  visible!:string
  countStar(star:any) {
    this.selectedValue = star;
    console.log('Value of star', this.selectedValue);
  }
  getOuterDivWidth():string{
    
    let w=this.selectedValue*28
    this.visible="hidden"
    console.log("outer",w)
    return w +'px'

  }
  getInnerDivWidth():string{

    return '150px'
  }
 
}
