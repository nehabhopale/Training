import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { RouteService } from '../routeService/route.service';
import { Passport } from '../models/passport';
import { User } from '../models/user';
import { Hobby } from '../models/hobby';
@Component({
  selector: 'app-user-register',
  templateUrl: './user-register.component.html',
  styleUrls: ['./user-register.component.css']
})
export class UserRegisterComponent implements OnInit {
  userForm!:FormGroup;
  hobby!:Hobby;
  status!:string;
  constructor(private routeService: RouteService,private router:Router) {
    this.userForm=new FormGroup({
      'firstName':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'lastName':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'email':new FormControl('',[Validators.required,Validators.email,Validators.maxLength(20)]),
      'password':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'country':new FormControl('',[Validators.required,Validators.maxLength(20)]),
      'hobbyName':new FormControl('',[Validators.required,Validators.maxLength(20)])
    });
   }


  ngOnInit(): void {
  }
  registerUser(){
    let newUser=new User();
    newUser.FirstName=this.userForm.get('firstName')?.value;
    newUser.LastName=this.userForm.get('lastName')?.value;
    newUser.Email=this.userForm.get('email')?.value;
    newUser.Password=this.userForm.get('password')?.value;
    newUser.Passport=new Passport();
    newUser.Passport.Country=this.userForm.get('country')?.value;
    newUser.Hobbies=[new Hobby()];
   
    for (var hobby of newUser.Hobbies){
      hobby.HobbyName =this.userForm.get('hobbyName')?.value;
    }
    // newUser.Hobbies.HobbyName=this.userForm.get('country')?.value;
   
    console.log(newUser);
    this.routeService.createUser(newUser).subscribe((data)=>{
     
      this.router.navigate(["/login"]);
    },(error)=>{
     console.log("Error in adding user",error);
    });
  }
  cancel(){
    this.status="";
  }

}
