import { Component } from '@angular/core';
import { stringify } from 'querystring';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'first-app';
  name:string;
 
  now: Date = new Date();
  constructor(){
    this.name="neha"
    console.log("instance created")
    setInterval(() => {
      this.now = new Date();
    }, 1000);
  }
  todaydate = new Date();
  a:number=0.234;
 
}
