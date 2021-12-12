package main

import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt")

type User struct {
	gorm.Model
	FullName string
	Address string
	IsMale *bool
}

func (U User)negate() {
	*U.IsMale = !*U.IsMale
}
// func CreateTable(db *gorm.DB){
// 	db.Table("users").DropTable(db)

// 	db.CreateTable(&User{})
// 	user := User{FullName: "Jinzhu",Address: "kop",IsMale:false}

// 	result := db.Create(&user) // pass pointer of data to Create

// 	fmt.Println(result.RowsAffected) // returns inserted records co
// }
func BoolAddr(b bool) *bool {
    boolVar := b
    return &boolVar
}

func updateRec(db *gorm.DB){
	db.Table("users").Debug().Where("FullName= ?", "Jinzgu").Update("IsMale", true)

}
// func Create(db *gorm.DB){

// 		user := User{FullName: "pteena", Address:"kop",IsMale:BoolAddr(true)}
// 		user.negate()
	
// 		db.NewRecord(user) // => returns `true` as primary key is blank
	
// 		db.Debug().Create(&user)
	
// 		db.NewRecord(user) // => return `false` after `user` creat
// }

// func deleteRecord(db *gorm.DB){
// 	db.Where("ID = ?", 1).Delete(&User{})
// }
func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/gomodel?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	fmt.Println(db)
	//CreateTable(db)
	//Create(db)
	//deleteRecord(db)
	updateRec(db)


}