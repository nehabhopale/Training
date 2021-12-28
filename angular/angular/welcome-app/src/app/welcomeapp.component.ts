import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './welcomeapp.component.html',
  styleUrls: ['./welcomeapp.component.css']
})
export class WelcomeAppComponent {
  title = 'welcome-app';
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
