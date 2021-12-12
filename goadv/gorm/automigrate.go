package main
import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt")

//AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column's type if its size, precision, nullable changed. It WON'T delete unused columns to protect your data.09-Nov-2021

type User struct {

	Name string
	ID int
	address string
}


func AutoMigrateTable(db *gorm.DB){

	db.AutoMigrate(&User{})
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
	AutoMigrateTable(db)
}