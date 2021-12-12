package main

import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt")

type Animal struct{
	Name string
}

func AutoMigrateTable(db *gorm.DB){

	db.AutoMigrate(&Animal{})
	user := Animal{Name: "tiger"}

	result := db.Create(&user) // pass pointer of data to Create

	fmt.Println(result.RowsAffected) // returns inserted records co
}

func CreateAnimals(db *gorm.DB) error {

	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()			//our transaction begins
  
	if err := tx.Error; err != nil {    //to print error occured while begining transactions
	  return err
	}
  
	if err := tx.Create(&Animal{Name: "Giraf"}).Error; err != nil {   
	   tx.Rollback()   				//it will rollback the operation
	   return err
	}

	//tx.SavePoint("sp1")
  
	if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
	   tx.Rollback()
	   return err
	}
	//tx.RollbackTo("sp1")
	return tx.Commit().Error
  }

func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/transaction?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	fmt.Println(db)
	AutoMigrateTable(db)
	CreateAnimals(db)
}
