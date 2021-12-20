package main

import (
"github.com/jinzhu/gorm"
"github.com/gorilla/mux"
"pass/model"
connector"pass/connector"
"fmt"
"log"
"net/http"
_"github.com/jinzhu/gorm/dialects/mysql")


func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/userApp?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	CreateTable(db)
	connector.Connect(db)
	route := mux.NewRouter()
	connector.RegisterUserRoutes(db,route ,route )
	connector.RegisterPassportRoutes(db,route )
	connector.RegisterHobbyRoutes(db,route )
	connector.RegisterCourseRoutes(db,route )
	log.Fatal(http.ListenAndServe(":9000", route ))
}

func CreateTable(db *gorm.DB){
	
	db.AutoMigrate(&model.Passport{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Hobby{})
	db.AutoMigrate(&model.Course{})
	
	err1 := db.Debug().Model(&model.Passport{}).AddForeignKey("uid","users(id)","CASCADE", "CASCADE").Error
	if err1 != nil {
			fmt.Println(err1)
	}
	err2:= db.Debug().Model(&model.Hobby{}).AddForeignKey("uid", "users(id)", "CASCADE", "CASCADE").Error
		if err2!= nil {
			fmt.Println(err2)
		}
}