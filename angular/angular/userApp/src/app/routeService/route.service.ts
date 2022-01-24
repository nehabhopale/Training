import { HttpClient,HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Course } from '../models/course';
import { Hobby } from '../models/hobby';
import { Passport } from '../models/passport';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class RouteService {
  varObserve!:Observable<any>
 
  constructor(private httpClient: HttpClient) { }
  getAllCourses():Observable<any>{
    let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
    return this.httpClient.get<any>("http://localhost:9000/courses",{headers:headers})
  }
  getUsers():Observable<any>{
    let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
    return this.httpClient.get<any>("http://localhost:9000/users",{headers:headers})
  }
  getHobby():Observable<any>{
    return this.httpClient.get<any>("http://localhost:9000/hobbies")
  }
  getPassport():Observable<any>{
    return this.httpClient.get<any>("http://localhost:9000/passports")
  }
  deletePassportById(id:string):Observable<any>{
    // let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
    return this.httpClient.delete<any>("http://localhost:9000/passports/"+id);
  }
  getPassportByUserId(userId:string):Observable<any>{
    
    return this.httpClient.get<any>("http://localhost:9000/users/"+userId+"/passports");
  }

  
  getHobbiesByUserId(userId:string):Observable<any>{
    return this.httpClient.get<any>("http://localhost:9000/hobbies/"+userId)

  }
  deleteCourseById(id:string):Observable<any>{
    return this.httpClient.delete<any>("http://localhost:9000/courses/"+id)

  }
  getUser(userId:string):Observable<any>{  
    return this.httpClient.get<any>("http://localhost:9000/users/"+userId);
  }

  
  
  addCourse(course:Course){
      // let currentToken:any=localStorage.getItem('Token')
      // let headers= new HttpHeaders().set('Token',currentToken);
     
       return this.httpClient.post<Course>("http://localhost:9000/courses",course) 
      //return this.http.post<any>("http://app:9000/courses",{"Name":courseName},{headers:headers}) 
    
  }
  updatePassport(passport:Passport){
    let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
    return this.httpClient.put<any>("http://localhost:9000/passports/"+passport.ID,passport,{headers:headers})
  }
  updateCourse(course:Course){
   let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
  return this.httpClient.put<any>("http://localhost:9000/courses/"+course.ID,course,{headers:headers})
  // return this.http.put<any>("http://app:9000/courses/"+id,{"Name":courseName},{headers:headers})
  }
  updateHobby(hobby:Hobby){
    console.log("in guarded route",localStorage.getItem('token'))
    let headers = new HttpHeaders().set('access_token',localStorage.getItem('token')!);
    return this.httpClient.put<any>("http://localhost:9000/hobbies/"+hobby.ID,hobby,{headers:headers})

  }
  createUser(user:User){

    console.log("inside add route")
    console.log("user inside route",user)
    
    return this.httpClient.post<User>("http://localhost:9000/users",user) 

  }
}
