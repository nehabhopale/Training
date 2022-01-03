import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-changeimage',
  templateUrl: './changeimage.component.html',
  styleUrls: ['./changeimage.component.css']
})
export class ChangeimageComponent implements OnInit {
  Hidespinner:boolean=true;
  buttonName:string;
  actualImage:string='./assets/n1.jfif'
  constructor() { 
    this.buttonName="hide"
  }

  ngOnInit(): void {
  }
  toggle(){
    this.Hidespinner=!this.Hidespinner
    if(this.Hidespinner){
      this.buttonName="show"
    }else{
      this.buttonName="hide"
    }
}
}