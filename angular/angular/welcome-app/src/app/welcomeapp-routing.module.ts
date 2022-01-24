import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RouteComponent } from './route/route.component';
// import { LoginComponent } from './login/login.component';

const routes: Routes = [
  {path:'dashboard/course',component:RouteComponent},
];
// this.router.navigate(["/dashboard/"+userId+"/passport"]);
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class WelcomeAppRoutingModule { }
