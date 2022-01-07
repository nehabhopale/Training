import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { interval } from 'rxjs';
import {  map } from 'rxjs/operators';
// import { pipe } from 'rxjs';

@Component({
  selector: 'app-observables',
  templateUrl: './observables.component.html',
  styleUrls: ['./observables.component.css']
})
export class ObservablesComponent implements OnInit {

  constructor() { }
  getAsync1():Observable<any>{
    const obs=new Observable<any>((observer)=>{
      observer.next(10)
      observer.next(20)
      observer.complete()
      observer.next(30) //code will not reach here
      
    })
    return obs
  }
  getAsync2(){
    const num=interval(2000)
    num.subscribe((data)=>{
      console.log(data)
    })
  }
  getAsync3(){
    return interval(2000)
      .pipe(map(n=>{
        return{time:new Date(),"counter":n}
      }))
    }

  ngOnInit(): void {
    this.getAsync1().subscribe((data)=>{
        console.log("data arrived")
        console.log(data)
      },(err)=>{
        console.log("in error")
        console.log(err)
      })
    // this.getAsync3().subscribe((data)=>{
    //   console.log("data arrived")
    //   console.log(data)
    // },(err)=>{
    //   console.log("in error")
    //   console.log(err)
    // })
    // this.getAsync2()
    console.log("on end of init")

  }
  
}
