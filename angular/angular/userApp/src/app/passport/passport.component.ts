import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { RouteService } from '../routeService/route.service';
export class Passport{
  ID?:string;
  CreatedBy?:string;
  CreatedAt?:string
  DeletedAt?:string;
  PassNo?:string;
  UID?:string;
  Country?:string;
}

@Component({
  selector: 'app-passport',
  templateUrl: './passport.component.html',
  styleUrls: ['./passport.component.css']
})
export class PassportComponent implements OnInit {
  observe!:Observable<any>
  passports!:Passport[]
  displayPassportModel="none"
  
  passportToBeUpdated!:Passport
  passportEditForm!:FormGroup;
  constructor(private routeService: RouteService) { 
    this.passportEditForm = new FormGroup({
      updatedPassCountryName : new FormControl('',[Validators.required,Validators.maxLength(20)])
      
    })
  }

  ngOnInit(): void {
    this.observe=this.routeService.getPassport()
    this.observe.subscribe((data)=>{
      console.log("data",data)
      this.passports=data
  })

  }
  updatePassportView(){

    this.observe=this.routeService.getPassport()
    this.observe.subscribe((data)=>{
      console.log(" DATA Passport FROM API ",data)
      console.log("data////",data)
      this.passports=data
    })
  } 
  deletePassport(passportId:string){
    this.routeService.deletePassportById(passportId).subscribe((data)=>{
      console.log("deleted passport");
      this.getPassports();
    },(error)=>{
     console.log("Error in deleting passport",error);
    })
  }
  updatePassport(myGroup:any){
    this.displayPassportModel="none"
     console.log("Updating hobby ",this.passportToBeUpdated.ID)
     console.log("New passport name ",myGroup.value.updatedPassCountryName)
  
     let updatePassport = new Passport()
     updatePassport.ID=this.passportToBeUpdated.ID
     updatePassport.Country=myGroup.value.updatedPassCountryName
   
     this.routeService.updatePassport(updatePassport).subscribe((data)=>{
       console.log("updating passport",data)
       this.updatePassportView()
     },(error)=>{
        
         console.log("Error updating passport name ",error)
       }
     )
   }
   openUpdatePassportModal(passport:Passport){
    this.displayPassportModel="block"
    this.passportToBeUpdated=passport
    console.log("passport updated id ",passport.ID)
  }
  closePassModal(){
    this.displayPassportModel="none"
  }
  getPassports(){
    let userId=localStorage.getItem('userId');
    this.routeService.getPassportByUserId(userId!).subscribe((data:any)=>{
      this.passports=data;
      console.log(data);
    });
  }

}
