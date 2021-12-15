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
	DNS:="root:admin@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	CreateTable(db)
	connector.Connect(db)
	
	r := mux.NewRouter()

	r.HandleFunc("/path/users", connector.GetUsers(db)).Methods("GET")
	r.HandleFunc("/path/addusers", connector.AddUser(db)).Methods("Post")
	r.HandleFunc("/path/users/{id}", connector.GetPassportByUserId(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func CreateTable(db *gorm.DB){
	
	db.AutoMigrate(&model.Passport{})
	db.AutoMigrate(&model.User{})
	
	err1 := db.Debug().Model(&model.Passport{}).AddForeignKey("uid","users(id)","CASCADE", "CASCADE").Error
	if err1 != nil {
			fmt.Println(err1)
	}
	
}