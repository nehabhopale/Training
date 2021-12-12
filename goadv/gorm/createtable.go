package main
import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt")

type User struct {

	Name string
	ID int
	address string
}


func CreateTable(db *gorm.DB){

	db.CreateTable(&User{})
	user := User{Name: "Jinzhu", ID: 2, address: "kop"}

	result := db.Create(&user) // pass pointer of data to Create

	fmt.Println(result.RowsAffected) // returns inserted records co
}
func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/godev?charset=utf8mb4&parseTime=True&loc=Local"
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