package main

import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt")


type User struct {

	Name string
	ID int
	address string
}


// func Create(db *gorm.DB){
// 	user := User{Name: "pooja", ID: 3, address:"kop"}

// 	db.NewRecord(user) // => returns `true` as primary key is blank

// 	db.Debug().Create(&user)

// 	db.NewRecord(user) // => return `false` after `user` creat
// }

func GetRecord(db *gorm.DB){
	var user User
	// Get first record, order by primary key
	db.Debug().First(&user)
	fmt.Println("first record",user)
	// Get one record, no specified order
	db.Debug().Take(&user)
	fmt.Println("random one  record",user)

	// Get last record, order by primary key
	db.Debug().Last(&user)
	fmt.Println("last record",user)
	var users []User
	// Get all records
	db.Debug().Find(&users)
	fmt.Println("all record",users)
	// Get record with primary key (only works for integer primary key)
	db.Debug().First(&user, 2)
	fmt.Println("specific record",user)
	
}
func update(db *gorm.DB){
	db.Table("users").Where("ID = ?", 3).Update("name", "rani")

}
func deleteRecord(db *gorm.DB){
	db.Where("ID = ?", 5).Delete(&User{})
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
	//fmt.Println(db)
	//Create(db)
	GetRecord(db)
	update(db)
	deleteRecord(db)
}
