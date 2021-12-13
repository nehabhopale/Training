package model
import (_"github.com/jinzhu/gorm/dialects/mysql"
uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm"
)

type User struct{
	CustomModel
	//ID uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	UserName string
	Hobbies []Hobby	`gorm:"foreignKey:UID"`
	Courses []Course `gorm:"association_autoupdate:false;association_autocreate:false;many2many:user_courses;"`
}
func NewUser( name string,hobbies []Hobby)*User{
	return &User{
		CustomModel:CustomModel{CreateBy:"neha"},
		UserName:name,
		Hobbies:hobbies,
		
	}
}
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
