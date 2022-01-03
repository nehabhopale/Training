import { Component, OnInit } from '@angular/core';
export interface Students {
  
  fname: string;
  lname:string;
}
@Component({
  selector: 'app-looping',
  templateUrl: './looping.component.html',
  styleUrls: ['./looping.component.css']
})

export class LoopingComponent implements OnInit {

  student!:Students[]
  constructor() { }

  ngOnInit(): void {
  }
  makeArray(size :number) :Array<any>{
    return new Array(size)

  }
  ArrayOfStudents() : Array<Students> {
    this.student = [
      { fname:"neha",
        lname: "bhopale"},
      { fname:"pooja",
        lname: "bhopale"},
      
    ];
    return this.student
  }


}
