import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import{RouteService} from '../routeService/route.service'

@Component({
  selector: 'app-route',
  templateUrl: './route.component.html',
  styleUrls: ['./route.component.css']
})
export class RouteComponent implements OnInit {
  observe!:Observable<any>
  constructor(private routeService:RouteService) { }
  courses:any[]=[]
  CName:any[]=[]
  ngOnInit(): void {
  //   this.observe=this.routeService.getAsync3()
  //  this.routeService.getAsync3().subscribe((data)=>{
  //     console.log("data arrived")
  //     console.log(data)
  //   },(err)=>{
  //     console.log("in error")
  //     console.log(err)
  //   })
  // }


  ////////////////////////////////
  // this.observe=this.routeService.getDataFromApi()
  //   this.observe.subscribe((data)=>{
  //     console.log("data",data)
  //     // console.log(data.)
  //     // console.log(" Data ",JSON.parse(data))
  //     // this.courses=JSON.parse(data)
  //     for(let c of data){
  //       this.CName.push(c.CourseName)
  //     }
  //     console.log("Course names ",this.CName)
  //   })

}

show(){
  this.observe=this.routeService.getDataFromApi()
    this.observe.subscribe((data)=>{
      console.log("data",data)
      // console.log(data.)
      // console.log(" Data ",JSON.parse(data))
      // this.courses=JSON.parse(data)
      for(let c of data){
        this.CName.push(c.CourseName)
      }
      console.log("Course names ",this.CName)
    })
  }
}
