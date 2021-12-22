package model

// import (uuid"github.com/satori/go.uuid"
// "time")

type User struct{
	Base
	//ID uuid.UUID	`gorm:";primary_key;type:varchar(50);"`
	FirstName string  
	LastName string
	Passport Passport	`gorm:"foreignKey:UID"`
	Email string		`gorm:"unique;not null"`
	Password string		`gorm:"not null"`
	Courses   []Course `gorm:"association_autoupdate:false;association_autocreate:false;many2many:user_courses;"`
	Hobbies   []Hobby	`gorm:"foreignKey:UID"`

}
