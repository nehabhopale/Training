package model

import (uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm")

type Passport struct{
		UID uuid.UUID	`gorm:"type:varchar(50)"`
		PassId uuid.UUID `gorm:"type:varchar(50);primary_key"`
}
func (pass *Passport) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("PassId", uuid.NewV4())
 }