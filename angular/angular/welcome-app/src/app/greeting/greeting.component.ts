import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-greeting',
  templateUrl: './greeting.component.html',
  styleUrls: ['./greeting.component.css']
})
export class GreetingComponent implements OnInit {
greeting!:string;
  constructor() { }

  ngOnInit(): void {
  }
  greet(name:string){
    var time = new Date();
    var currentTime=time.getHours();

    if (6 <= currentTime && currentTime < 12 ){
      this.greeting="hi"+" "+name+" "+"good morning"
    } else if (12 <= currentTime && currentTime < 16 ){
      this.greeting="hi"+" "+name+" "+"good afternoon"
    } else if (16 <= currentTime && currentTime <= 20 ){
      this.greeting="hi"+" "+name+" "+"good evening"
    } else {
      this.greeting="hi"+" "+name+" "+"good night"
    }
      
      // if (input < 12) {
      //   return 'Good Morning';
      // } else if (input >= 12 && input <= 17) {
      //   return 'Good Afternoon';
      // } else if (input > 17 && input <= 24) {
      //   return 'Good Evening';
      // } else {
      //   return "I'm not sure what time it is!";
      // }
  }
}
