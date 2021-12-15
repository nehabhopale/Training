package model

import (uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm")

type User struct{
	ID uuid.UUID	`gorm:";primary_key;type:varchar(50);"`
	UserName string 
	Passport Passport	`gorm:"foreignKey:UID"`

}
func NewUser( name string)User{
	return User{
		Passport:Passport{PassId:uuid.NewV4()},
		UserName:name,
		
		
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
