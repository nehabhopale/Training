import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrls: ['./parent.component.css']
})
export class ParentComponent implements OnInit {
  rating!: number;
  constructor() { }

  ngOnInit(): void {
  }
  // countStar(star:any) {
  //   this.rating = star;
  //   console.log('Value of star', this.rating);
  // }
}
