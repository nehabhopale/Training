package main
import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	UserName string
	Languages []Language `gorm:"many2many:user_languages;"`
  }
  
type Language struct {
	gorm.Model
	Name string
	Users []User `gorm:"many2many:user_languages;"`
}
func CreateTable(db *gorm.DB){
		//db.Table("users").DropTable(db)
		db.AutoMigrate(&User{})  
		db.AutoMigrate(&Language{})
		user:=User{
			UserName :"neha",
			Languages:[]Language{
				{Name:"english"},
				{Name:"marathi"},
			},
		}
		db.Debug().Create(&user)

		var user11 User
		db.Model(&user11).Association("Languages")
 }

 func Association(db *gorm.DB){
	var user User
	db.Model(&user).Association("Languages")
	db.Model(&user).Association("Languages").Append(&Language{Name: "hindi"})
 }

 func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/manykey?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	fmt.Println(db)
	
	CreateTable(db)
	Association(db)
}