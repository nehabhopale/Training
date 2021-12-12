package main
import (
	"github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt"
)
type User struct {
	gorm.Model
	Name       string 
	CreditCard CreditCard `gorm:"foreignKey:UserName;references:name;Column:CreditCard"`
	// use UserName as foreign key
}
  
type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}

func CreateTable(db *gorm.DB){
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})
	//foreign key username is from table creditcard
	
	user1:=	User{
		Name:"neha",
		CreditCard:CreditCard{Number:"456234",UserName:"neha"},
		//Hobby:Hobby{StuId:1,Hob:"mri"},
		
	}
	db.Create(&user1)
	db.Model(&CreditCard{}).AddForeignKey("user_name","users(credit_card)","RESTRICT", "RESTRICT")
	// card:=CreditCard{Number:"456234",UserName:"neha"}

	// db.Create(&card)
}
func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/relation?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	fmt.Println(db)
	
	CreateTable(db)
}