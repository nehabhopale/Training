import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { LoginserviceService } from '../loginservice/loginservice.service';
import { RouteService } from '../routeService/route.service';

@Injectable({
  providedIn: 'root'
})
export class GuardService implements CanActivate {

  constructor(private myService:LoginserviceService,public router: Router) { }
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    console.log("In auth service");
    let token=localStorage.getItem('token');
    console.log("token getting********",token)
    if(!token){
      this.router.navigate(['login']);
      return false;
    }
    this.myService.checkToken(token).subscribe((data)=>{
      console.log(data);
      
      return true;
    },(error)=>{
      console.log(error)
      this.router.navigate(['login']);
      return false;
    });
    return true;
  }
  
}
