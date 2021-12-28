import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-test',
  templateUrl: './mytest.component.html',
  styleUrls: ['./mytest.component.css']
})
export class TestComponent implements OnInit {
  condition:boolean;
  name:string;
  a:number;
  showtime!:boolean;
  now: Date = new Date();
  constructor() { 
    this.condition=true
    this.name="xyz"
    this.a=3
    setInterval(() => {
      this.now = new Date();
    }, 1000);
  }
  showMe(){
    this.showtime=true
  }
  hideMe(){
    this.showtime=false
  }

  ngOnInit(): void {
    
  }

}
