package model

import (uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm")

type User struct{
	Base
	ID uuid.UUID	`gorm:";primary_key;type:varchar(50);"`
	FirstName string 
	LastName string
	Passport Passport	`gorm:"foreignKey:UID"`
	Email string		`gorm:"unique;not null"`
	Password string
	Courses   []Course `gorm:"association_autoupdate:false;association_autocreate:false;many2many:user_courses;"`
	Hobbies   []Hobby

}
func NewUser(fname string,lname string,email string,pass string)User{
	return User{
		Passport:Passport{PassId:uuid.NewV4()},
		FirstName:fname,
		LastName:lname,
		Email:email,
		Password:pass,	
	}
}
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
// type Passport struct{
// 	UID int
// 	PassId int //`gorm:";primary_key"`
// }
// type Model struct{
// 	ID int		`gorm:";primary_key"`
// 	CreateBy  string
// }
