package main

import (
	"fmt"

	//"test/model"
	repo"test/repo"
	"test/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
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
	
		//Create Table and Adding Foreign Key
	// err1 := db.Debug().CreateTable(&model.User{}, &model.Course{}, &model.Hobby{}).Error
	// if err1 != nil {
	// 		fmt.Println(err)
	// }
	// err2 := db.Debug().Model(&model.Hobby{}).AddForeignKey("uid","users(id)","RESTRICT", "RESTRICT")
	// if err2 != nil {
	// 		fmt.Println(err2)
	// }

	unit1 := repo.NewUnitOfWork(db, true)
	repo1 := repo.NewRepository()
	fmt.Println(unit1, repo1)

	userServices:=services.NewUser(repo1,unit1)

	//userServices.CreateUser()
	//userServices.GetAllUser()
	userServices.GetUser()

}