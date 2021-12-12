package main

import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt"
uuid"github.com/satori/go.uuid"
"test/ex")

type User struct{
	Id uuid.UUID	`gorm:"type:varchar(36);primaryKey`
	UserName string
	Hobbies []Hobby	`gorm:"foreignKey:UserId"`
	courses []Course `gorm:"many2many:user_languages;"`
}

//user should be able to access courses 
type Course struct{
	users []User
	CourseName string
	CourseId uuid.UUID		`gorm:"type:varchar(36);primaryKey`
}
//from users we can add hobbies 
type Hobby struct{
	UserId uuid.UUID 
	HobbyId  uuid.UUID		`gorm:"primaryKey"`
	HobbyName string
}
func CreateTable(db *gorm.DB){
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Hobby{})
}
func main(){
	repo:=ex.NewRepository()
	DNS:="root:admin@tcp(127.0.0.1:3306)/repo?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	uow:=ex.NewUnitOfWork(db,true)
	CreateTable(db)
	var user User
	c:=[]string{"hobbies","courses"}
	err1:=repo.Get(uow , &user, user.Id, c)
	fmt.Println(err1)
}