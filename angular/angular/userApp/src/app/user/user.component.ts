import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../models/user';
import { RouteService } from '../routeService/route.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  observe!:Observable<any>
  users!:User[]

  constructor(private routeService: RouteService) { }

  ngOnInit(): void {
    this.observe=this.routeService.getUsers()
    this.observe.subscribe((data)=>{
      console.log("data",data)
      this.users=data
  })
  }


}
