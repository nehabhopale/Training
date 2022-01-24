import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { RouteService } from '../routeService/route.service';
import { Hobby } from '../models/hobby';
// export class Hobby{
//   ID?:string;
//   CreatedBy?:string;
//   CreatedAt?:string
//   DeletedAt?:string;
//   HobbyName?:string;
//   UID?:string;
// }

@Component({
  selector: 'app-hobby',
  templateUrl: './hobby.component.html',
  styleUrls: ['./hobby.component.css']
})
export class HobbyComponent implements OnInit {
  observe!:Observable<any>
  hobbies!:Hobby[]
  displayHobbyModel="none"
  hobbyToBeUpdated!:Hobby
  hobbyEditForm:FormGroup;
  editData:any={mode:"",hobby:Hobby};
  constructor(private routeService: RouteService) {
    this.hobbyEditForm = new FormGroup({
      updatedHobbyName : new FormControl('',[Validators.required,Validators.maxLength(50)])
      
    })
   }

  ngOnInit(): void {
    this.observe=this.routeService.getHobby()
    this.observe.subscribe((data)=>{
      console.log("data",data)
      this.hobbies=data
  })
  }
  
  updateHobbyView(){

    this.observe=this.routeService.getHobby()
    this.observe.subscribe((data)=>{
      console.log(" DATA Hobby FROM API ",data)
      console.log("data////",data)
      this.hobbies=data
    })
  } 
  getHobbies(){
    let userId=localStorage.getItem('userId');
    this.routeService.getHobbiesByUserId(userId!).subscribe((data:any)=>{
      this.hobbies=data;
      console.log(data);
    });
  }
  updateHobby(myGroup:any){
    this.displayHobbyModel="none"
     console.log("Updating hobby ",this.hobbyToBeUpdated.ID)
     console.log("New hobby name ",myGroup.value.updatedHobbyName)
  
     let updateHobby = new Hobby()
     updateHobby.ID=this.hobbyToBeUpdated.ID
     updateHobby.HobbyName=myGroup.value.updatedHobbyName
   
     this.routeService.updateHobby(updateHobby).subscribe((data)=>{
       console.log("updating hobby",data)
       this.updateHobbyView()
     },(error)=>{
        
         console.log("Error updating hobby name ",error)
       }
     )
   }
   openUpdateHobbyModal(hobby:Hobby){
    this.displayHobbyModel="block"
    this.hobbyToBeUpdated=hobby
    console.log("Hobby updated id ",hobby.ID)
  }
  closeHobbyModal(){
    this.displayHobbyModel="none"
  }

}


