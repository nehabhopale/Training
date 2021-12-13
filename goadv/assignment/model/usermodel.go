package model
import (_"github.com/jinzhu/gorm/dialects/mysql"
uuid"github.com/satori/go.uuid"
)

type User struct{
	ID uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	UserName string
	Hobbies []Hobby	`gorm:"foreignKey:UID"`
	Courses []Course `gorm:"association_autoupdate:false;association_autocreate:false;many2many:user_courses;"`
}
func NewUser(ID uuid.UUID, name string,hobbies []Hobby)*User{
	return &User{
		ID:ID,
		UserName:name,
		Hobbies:hobbies,
		
	}
}
// func (u *User) AddHobbies(h Hobby) {
// 	u.Hobbies = append(u.Hobbies, h)
// }