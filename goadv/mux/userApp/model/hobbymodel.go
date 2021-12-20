package model
import (_"github.com/jinzhu/gorm/dialects/mysql"
uuid"github.com/satori/go.uuid")
// "time"
// "github.com/jinzhu/gorm")
//from users we can add hobbies 
type Hobby struct{
	Base
	UID uuid.UUID `gorm:"type:varchar(36);"`
	//HobbyID  uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	HobbyName string
}
// func NewHobby(Name string)*Hobby{
// 	return &Hobby{
// 		Base:Base{CreateBy:"neha",ID:uuid.NewV4(),CreateAt:time.Now()},
// 	}
// }
// func (hobby *Hobby) BeforeCreate(scope *gorm.Scope) error {
	
// 	scope.SetColumn("ID", uuid.NewV4())
// 	scope.SetColumn("CreateBy", "neha")
// 	scope.SetColumn("CreateAt", time.Now())
// 	return nil
// }