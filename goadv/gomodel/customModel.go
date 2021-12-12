package main
import ("github.com/jinzhu/gorm"
_"github.com/jinzhu/gorm/dialects/mysql"
"fmt"
uuid"github.com/satori/go.uuid")
//"github.com/google/uuid")

type People struct{
	Model
	FirstName string
	Age int 
	IsMale bool

}

type  Model struct{
	
	ID uuid.UUID `gorm:"type:varchar(36);primaryKey`
	Email string
}

func table(db *gorm.DB){
	db.Table("peoples").DropTable(db)
	db.Table("models").DropTable(db)
	db.CreateTable(&People{})
	db.CreateTable(&Model{})			//if only model table created then no enteries will be added
	 user := People{
			Model:Model{ID:uuid.NewV1(),Email:"neha@"},
			FirstName: "Jinzhu",
			Age:2,
			IsMale:false,
		}
	
		result := db.Create(&user) // pass pointer of data to Create
	
		 fmt.Println(result.RowsAffected) // returns inserted records co


}

func CreateTable(db *gorm.DB){
	 	db.Table("peoples").DropTable(db)
		db.Table("models").DropTable(db)
		db.AutoMigrate(&People{})  			//only people (parent)is automigrated automigrate(creates table people with own and model fields)
		db.AutoMigrate(&Model{})			//when both are automigated then people and model tables created with model as empty table
		
		user := People{
			Model:Model{ID:uuid.NewV1(),Email:"neha@"},
			FirstName: "Jinzhu",
			Age:2,
			IsMale:false,
		}
	
		result := db.Create(&user) // pass pointer of data to Create
	
		 fmt.Println(result.RowsAffected) // returns inserted records co
}

func CreateRecord(db *gorm.DB){
	user := People{
		Model:Model{ID:uuid.NewV1(),Email:"bneha@"},
		FirstName: "Jin",
		Age:8,
		IsMale:true,
	}

	db.NewRecord(user) // => returns `true` as primary key is blank

	db.Debug().Create(&user)

	db.NewRecord(user) // => return `false` after `user` creat
}
func GetRecord(db *gorm.DB){
	var user People
	// Get first record, order by primary key
	db.Debug().First(&user)
	fmt.Println("first record",user)
	// Get one record, no specified order
	db.Debug().Take(&user)
	fmt.Println("random one  record",user)

	// Get last record, order by primary key
	db.Debug().Last(&user)
	fmt.Println("last record",user)
	var users []People
	// Get all records
	db.Debug().Find(&users)
	fmt.Println("all record",users)
	
}
func update(db *gorm.DB){
	var people People
	people.IsMale=false
	db.Model(&people).Debug().Update(&people)

}
func updateWithInterface(db *gorm.DB){
	var people People

	db.Model(&people).Updates(map[string]interface{}{"FirstName": "Jinzhu", "age": 2, "IsMale":true })
// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
}
func deleteRecord(db *gorm.DB){
	db.Where("email = ?", "bneha@").Delete(&People{})
}


func main(){
	DNS:="root:admin@tcp(127.0.0.1:3306)/gocustom?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open("mysql",DNS)
	if err!=nil{
		fmt.Println(err)
		panic("connection failed")
	}else{
		fmt.Println("connected to db")
	}
	fmt.Println(db)
	
	CreateTable(db)
	
	GetRecord(db)
	update(db)
	deleteRecord(db)
	updateWithInterface(db)

	table(db)
	CreateRecord(db)
}