import { Injectable } from '@angular/core';
import { interval, Observable } from 'rxjs';
import {  map } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class RouteService {
  varObserve!:Observable<any>
  serviceBaseUrl!:'http://localhost:9000/courses'
  constructor(private httpClient: HttpClient) { }
  getAsync3():Observable<any>{
    this.varObserve=interval(5000)
      .pipe(map(n=>{
        return{"time":new Date(),"counter":this.getRandomInt(1,36)}
      }))
    return this.varObserve

  }
  getRandomInt(min:number, max:number) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }
  // getRandomInt() {
    
  //   console.log("client data",this.httpClient.get(this.serviceBaseUrl))
  // }
  getDataFromApi():Observable<any>{
    console.log("inside get data")
    console.log(this.serviceBaseUrl)
    return this.httpClient.get("http://localhost:9000/courses")
}}
 


