package model
import (_"github.com/jinzhu/gorm/dialects/mysql"
uuid"github.com/satori/go.uuid"
)
//from users we can add hobbies 
type Hobby struct{
	UID uuid.UUID `gorm:"type:varchar(50);"`
	HobbyID  uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	HobbyName string
}
// func NewHobby(HobbyID uuid.UUID,Name string)*Hobby{
// 	return &Hobby{
// 		HobbyID:HobbyID,
// 		HobbyName:Name,
// 	}

// }