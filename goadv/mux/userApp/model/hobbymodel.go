package model
import (_"github.com/jinzhu/gorm/dialects/mysql"
uuid"github.com/satori/go.uuid")

//from users we can add hobbies 
type Hobby struct{
	Base
	UID uuid.UUID `gorm:"type:varchar(36);"`
	//HobbyID  uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	HobbyName string
}
