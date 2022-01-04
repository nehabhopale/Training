import { Component, OnInit } from '@angular/core';
export interface Students {
  id:number,
  cgpa:number,
  name: string,
  dob:string
}
@Component({
  selector: 'app-students',
  templateUrl: './students.component.html',
  styleUrls: ['./students.component.css']
})
export class StudentsComponent implements OnInit {
  student!:Students[]
  BackColor!:string;
  constructor() { }

  ngOnInit(): void {
  }
  ArrayOfStudents() : Array<Students> {
    this.student = [
      { id:1,
        cgpa:9,
        name:"neha",
        dob:"26-8-1999"
      },
      {id:2,
       cgpa:7.3,
       name: "pooja",
       dob:"21-3-1997"
      },
      {
      id:3,
      cgpa:5,
      name:"sweety",
      dob:"24-3-1996"
      },
      {id:4,
      cgpa:8,
      name:"rani",
      dob:"12-9-1998"
      }
      
    ];
    return this.student
  }
  
}
