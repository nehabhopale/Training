import { Injectable } from '@angular/core';
import { interval, map, Observable } from 'rxjs';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class LoginserviceService {

  constructor(private http:HttpClient) { }
  getRandomInt(max:number) {
    return Math.floor(Math.random() * max);
  }

  

  getCourses():Observable<any>{
    return this.http.get<any>("http://localhost:9000/courses/");
  }

 

  login(email:string,password:string):Observable<any>{
    return this.http.post<any>("http://localhost:9000/login",{"email":email,"password":password});
  }

  checkToken(token:string):Observable<any>{
    return this.http.post<any>("http://localhost:9000/login/checktoken",{"token":token});
  }
}
