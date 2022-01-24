import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { CourseComponent } from './course/course.component';
import { HttpClientModule } from '@angular/common/http';
import { HobbyComponent } from './hobby/hobby.component';
import { PassportComponent } from './passport/passport.component';
import { UserRegisterComponent } from './user-register/user-register.component';
import { UserComponent } from './user/user.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DashboardComponent,
    CourseComponent,
    HobbyComponent,
    PassportComponent,
    UserRegisterComponent,
    UserComponent
  ],
  imports: [
    BrowserModule,FormsModule,ReactiveFormsModule,
    AppRoutingModule,HttpClientModule,NgbModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
