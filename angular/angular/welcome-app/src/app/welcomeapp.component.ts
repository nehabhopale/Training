import { Component } from '@angular/core';
//import{CheckerDirective}from './directives/checker.directive'

@Component({
  selector: 'app-root',
  templateUrl: './welcomeapp.component.html',
  styleUrls: ['./welcomeapp.component.css'],
  //directives: [CheckerDirective]
})
export class WelcomeAppComponent {
  title = 'welcome-app';
  name:string;
  condition:boolean;
 
 now: Date = new Date();
  constructor(){
    this.name="neha"
    this.condition=false
    console.log("instance created")
    setInterval(() => {
      this.now = new Date();
    }, 1000);
  }
  todaydate = new Date();
  a:number=0.234;
}
