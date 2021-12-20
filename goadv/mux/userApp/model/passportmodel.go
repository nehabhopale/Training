package model

import (uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm"
"time"
)

type Passport struct{
		Base
		UID uuid.UUID	`gorm:"type:varchar(50)"`
		PassNo uuid.UUID `gorm:"type:varchar(50)"`
		Country string
}


func (pass *Passport) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("PassNo", uuid.NewV4())
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreateBy", "neha")
	scope.SetColumn("CreateAt", time.Now())
	return nil
}