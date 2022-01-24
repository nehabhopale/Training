import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { flatMap, Observable } from 'rxjs';
import { RouteService } from '../routeService/route.service';
import { Course } from '../models/course';
// export class Course{
//   ID?:string;
//   CreatedBy?:string;
//   CreatedAt?:string;
//   DeletedAt?:string;
//   CourseName?:string;
//   Prize?:number;
// }
@Component({
  selector: 'app-course',
  templateUrl: './course.component.html',
  styleUrls: ['./course.component.css']
})

export class CourseComponent implements OnInit {
  observe!:Observable<any>
  courses!:Course[]
  courseName!:string
  courseToBeUpdated!:Course
  // prizeToBeUpdate!:Course
  prize!:number
  myGroup!:FormGroup
  display="none"
  displayCourseModel="none"
  constructor(private routeService: RouteService) {
    this.myGroup = new FormGroup({
      updatedCourseName : new FormControl('',[Validators.required,Validators.maxLength(50)]),
      updatedprize: new FormControl('',[Validators.required,Validators.maxLength(30)])
    })
   }
  
  // CName:any[]=[]
  ngOnInit(): void {
    this.observe=this.routeService.getAllCourses()
    this.observe.subscribe((data)=>{
      console.log("data",data)
      this.courses=data
    })
  }
  deleteCourse(courseId:string){
      this.routeService.deleteCourseById(courseId).subscribe((data)=>{
      // this.message={text:"successfully deleted course",status:"success"};
      this.getCourses();
    },
    (error)=>{
    console.log("Error in deleting course",error);
    });
  }
  getCourses(){
    let userId=localStorage.getItem('userId');
    this.routeService.getUser(userId!).subscribe((data)=>{
      this.courses=data.Courses!;
      console.log("}}}}}}}}}}",this.courses);
    },(error)=>{
      console.log("Error in getting courses",error);
    })
  }
  updateCourseView(){

    this.observe=this.routeService.getAllCourses()
    this.observe.subscribe((data)=>{
      console.log(" DATA COURSE FROM API ",data)
      console.log("data////",data)

      this.courses=data
    })
  }
    
  
  addCourse(){
    console.log("New course name ",this.courseName)
    let newCourse = new Course()
    newCourse.CourseName=this.courseName
    newCourse.Prize=this.prize
    console.log("newCourse",newCourse)
    this.routeService.addCourse(newCourse).subscribe((data)=>{
        console.log("adding course",data)
        this.updateCourseView()
      },(error)=>{
        
        console.log("Error adding new course ",error)
      })
      this.display="none"
  }
  closeCourseModal(){
    this.displayCourseModel="none"
  }
  updateCourse(myGroup:any){
   this.displayCourseModel="none"
    console.log("Updating course ",this.courseToBeUpdated.ID)
    console.log("New course name ",myGroup.value.updatedCourseName)
    console.log("new prize",myGroup.value.updatedprize)
    let updateCourse = new Course()
    updateCourse.ID=this.courseToBeUpdated.ID
    updateCourse.CourseName=myGroup.value.updatedCourseName
    updateCourse.Prize=myGroup.value.updatedprize
    this.routeService.updateCourse(updateCourse).subscribe((data)=>{
      console.log("adding course",data)
      this.updateCourseView()
    },(error)=>{
       
        console.log("Error updating course name ",error)
      }
    )
  }
  onCloseHandled(){
    this.display="none"
  }
  openModal(){
    this.display="block"
  }
  openUpdateCourseModal(course:Course){
    this.displayCourseModel="block"
    this.courseToBeUpdated=course
    console.log("Course updated id ",course.ID)
  }
}
  // show(){
    // this.observe=this.courseService.getDataFromApi()
    //   this.observe.subscribe((data)=>{
    //     console.log("data",data)
    //     // console.log(data.)
    //     // console.log(" Data ",JSON.parse(data))
    //     // this.courses=JSON.parse(data)
    //     for(let c of data){
    //       // this.CName.push(c.CourseName)
    //       this.courses.push(c)
    //     }
    //     // console.log("Course names ",this.CName)
    //   })
    // }


