package main
import (
	"github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt"
//uuid"github.com/satori/go.uuid"
)


//has many relation
type Student struct{
	Hobbies []Hobby	 `gorm:"foreignKey:StuId"`
	ID int	`gorm:"primaryKey"`//`gorm:"column:foreign_key;"`
	Name string
}
type Hobby struct{
	Hid int		`gorm:"primaryKey;Column:HobbyPrimaryKey"`
	StuId int   
	Hob string
}

func CreateTable(db *gorm.DB){
	db.Table("hobbies").DropTable(db)
	db.Table("students").DropTable(db)
 	db.AutoMigrate(&Student{})  
	db.AutoMigrate(&Hobby{})  
	db.Model(&Hobby{}).AddForeignKey("stu_id","students(id)","RESTRICT", "RESTRICT")
	user1:=	Student{
		//Hobby:Hobby{StuId:1,Hob:"mri"},
		ID:1,
		Name:"neha",
	}
	db.Create(&user1)
	hob1:=Hobby{Hid:2,StuId:1,Hob:"cricket"}

	db.Create(&hob1)

	hob2:=Hobby{Hid:3,StuId:5,Hob:"sport"}		//if we do stuid5 error (adding fk whose pk not exists)
	db.Create(&hob2)
	
	
}

func preload(db *gorm.DB){
	var students []Student
	//db.Debug().Find(&students)	// SELECT * FROM `students`
    f:=db.Debug().Preload("Hobbies").Find(&students).Error
	fmt.Println(f)

    // fmt.Println(students)
	// for _,user := range students {
	// 	db.Debug().Where("ID", user.ID).Find(&user.Hobbies)
	//   }
}


// }
// // func CreateRecord(db *gorm.DB){
	
// // 	user:=Student{
// // 		Hobby:Hobby{StuId:1,Hob:"tri"},
// // 		ID:2,
// // 		Name:"pooja",

// // 	}
// // 	db.NewRecord(user) // => returns `true` as primary key is blank

// // 	db.Debug().Create(&user)
// // 	db.NewRecord(user) // => return `false` after `user` creat

// }
func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/fkey?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	fmt.Println(db)
	
	CreateTable(db)
	preload(db)
	// sCreateRecord(db)
}