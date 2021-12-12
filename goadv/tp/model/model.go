package model

import (_"github.com/jinzhu/gorm/dialects/mysql"
uuid"github.com/satori/go.uuid"
)

type User struct{
	ID uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	UserName string
	Hobbies []Hobby	`gorm:"foreignKey:UserID"`
	Courses []Course `gorm:"association_autoupdate:false;association_autocreate:false;many2many:user_courses;"`
}

//user should be able to access courses 
type Course struct{
	users []User `gorm:"many2many:user_courses;"`
	CourseName string
	CourseID uuid.UUID `gorm:"primary_key;type:varchar(50);"`
}
//from users we can add hobbies 
type Hobby struct{
	UID uuid.UUID `gorm:"type:varchar(50);"`
	HobbyID  uuid.UUID `gorm:"primary_key;type:varchar(50);"`
	HobbyName string
}