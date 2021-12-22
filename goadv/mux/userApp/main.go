package main

import (
"github.com/jinzhu/gorm"
"github.com/gorilla/mux"
"pass/model"
"os"
connector"pass/connector"
repo"pass/repository"
services"pass/services"
"fmt"
"log"
"net/http"
_"github.com/jinzhu/gorm/dialects/mysql"
"pass/handler"
"github.com/rs/zerolog"
"io/ioutil")


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

	repo1 := repo.NewRepository()
	tempFile,_:= ioutil.TempFile(os.TempDir(), "deleteme")
	logger := zerolog.New(tempFile).With().Logger()

	userService:=services.NewUserService(repo1,&logger,db)
	passportService:=services.NewPassportService(repo1,&logger,db)
	hobbyService:=services.NewHobbyService(repo1,&logger,db)
	courseService:=services.NewCourseService(repo1,&logger,db)
	handler:=handler.Newhandler(userService)


	userConnector:=connector.NewUserConnector(handler,userService,passportService)
	passportConnector:=connector.NewPassportConnector(handler,userService,passportService)
	hobbyConnector:=connector.NewHobbyConnector(handler,userService,hobbyService)
	courseConnector:=connector.NewCourseConnector(handler,userService,courseService)

	router := mux.NewRouter()
	router.HandleFunc("/login", handler.GetTokenHandler).Methods("POST")
	authRoute:=router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true
	}).Subrouter()
	nonAuthRoute:=router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return true 
	 }).Subrouter()

	userConnector.RegisterUserRoutes(authRoute ,nonAuthRoute)
	passportConnector.RegisterPassportRoutes(authRoute ,nonAuthRoute)
	hobbyConnector.RegisterHobbyRoutes(authRoute ,nonAuthRoute)
	courseConnector.RegisterCourseRoutes(authRoute ,nonAuthRoute)
	log.Fatal(http.ListenAndServe(":9000", router))
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