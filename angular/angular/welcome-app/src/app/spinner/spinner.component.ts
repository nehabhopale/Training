import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-spinner',
  templateUrl: './spinner.component.html',
  styleUrls: ['./spinner.component.css']
})
export class SpinnerComponent implements OnInit {

  images:string[]=['./assets/n1.jfif','./assets/n2.jfif'];
  actualImage: string;
  Counter = 0;
 

  constructor() {
    // this.imageUrl:this.images[0]
    this.actualImage = this.images[0];
    setInterval(() => {
      this.Counter++;
      if (this.Counter > this.images.length - 1) {
        this.Counter = 0;
      }
      this.actualImage = this.images[this.Counter];
    }, 2000);
   }
   

  ngOnInit(): void {
   
    
  }

}
