import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CourseComponent } from './course/course.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { GuardService } from './guard/guard.service';
import { HobbyComponent } from './hobby/hobby.component';
import { LoginComponent } from './login/login.component';
import { PassportComponent } from './passport/passport.component';
import { UserRegisterComponent } from './user-register/user-register.component';
import { UserComponent } from './user/user.component';

const routes: Routes = [
  {path:'',redirectTo:'login',pathMatch:'full'},
  {path:'login',component:LoginComponent},
  {path:'register',component: UserRegisterComponent},
  {path:'dashboard',component:DashboardComponent,canActivate:[GuardService],children:[
    {path:'course',component:CourseComponent},
    {path:'hobby',component:HobbyComponent},
    {path:'passport',component:PassportComponent},
    {path:'user',component:UserComponent}
  ]},
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
