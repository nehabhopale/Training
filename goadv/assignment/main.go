package main

import (
	"fmt"
	"test/model"
	repo"test/repo"
	 "test/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid"github.com/satori/go.uuid"
	
)

var db *gorm.DB
func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/repo1?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
			fmt.Println(err)
			panic("connection failed")
	}else{
			fmt.Println("connected to db")
	}
	CreateTable(db)
	repository := repo.NewRepository()
	userServices:=services.NewUserService(repository)
	var hobbies = []model.Hobby{{CustomModel:model.CustomModel{ID: uuid.NewV4()}, HobbyName: "sports"},
								{CustomModel:model.CustomModel{ID: uuid.NewV4()}, HobbyName: "cooking"},
							    }
			//add users
	user:=model.NewUser("pooja",hobbies)
	
	userServices.AddUser(db,user)
	userServices.GetUser(db)
	//get  all users
	var users []model.User
	var str1 = []string{ "Courses","Hobbies"}
	userServices.GetUsers(db,&users, str1)
	fmt.Println(users)

	// //get user from id 
	// var user model.User

	// Id1,_:= uuid.FromString("2ce504bc-03bc-4610-9a9d-60e97d831926")
	// userServices.GetUserFromId(db,&user,Id1,str1)

	// fmt.Println(user)

	// //update user
	id2, _ := uuid.FromString("2ce504bc-03bc-4610-9a9d-60e97d831926")
	var userToBeUpdated model.User
	var str = []string{"Courses", "Hobbies"}
	userServices.GetUserFromId(db,&userToBeUpdated, id2, str)
	userToBeUpdated.UserName = "neha"
	userServices.UpdateUser(db,userToBeUpdated)	
	fmt.Println(userToBeUpdated)


	//delete user 
	id3, _ := uuid.FromString("2ce504bc-03bc-4610-9a9d-60e97d831926")
	var userToBeDeleted model.User
	var str2 = []string{"Courses", "Hobbies"}
	userServices.GetUserFromId(db,&userToBeDeleted, id3, str2)
	userServices.DeleteUser(db,userToBeDeleted)
	fmt.Println(userToBeDeleted)
	
	//**********************************************************************//

	courseServices:=services.NewCourseService(repository)
	//add courses
	course:=model.NewCourse("golang")

	courseServices.AddCourse(db,course)
	err2:=db.Debug().Model(course).Association("Users").Error
	if err2!=nil{
		fmt.Println("error in association------>",err2)
	}


	//get  all coursess
	var courses []model.Course
	 var str11 []string
	courseServices.GetCourses(db,&courses, str11)
	fmt.Println(courses)

	//get course from id 
	 var course1 model.Course

	 Id11,_:= uuid.FromString("b18a7c11-ed9f-4172-88ea-f620bf2e8171")
	 courseServices.GetCourseFromId(db,&course1,Id11,str11)

	 fmt.Println(course1)

	// //upddate course

	var courseToBeUpdated model.Course
	id22, _ := uuid.FromString("b18a7c11-ed9f-4172-88ea-f620bf2e8171")
	courseServices.GetCourseFromId(db,&courseToBeUpdated, id22, str11)
	courseToBeUpdated.CourseName = "golang"
	courseServices.UpdateCourse(db,courseToBeUpdated)	
	fmt.Println(courseToBeUpdated)


 	// //delete course
	id33, _ := uuid.FromString("450c14bf-ddae-45d4-a03a-4a92a3e4bb14")
	var courseToBeDeleted model.Course
	courseServices.GetCourseFromId(db,&courseToBeDeleted, id33, str11)
	courseServices.DeleteCourse(db,courseToBeDeleted)
	fmt.Println(courseToBeDeleted)

}
func CreateTable(db *gorm.DB){
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Hobby{})
	db.AutoMigrate(&model.Course{})

	err1 := db.Debug().Model(&model.Hobby{}).AddForeignKey("uid","users(id)","CASCADE", "CASCADE").Error
	if err1 != nil {
			fmt.Println(err1)
	}
	
}