package model

import (uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm"
"time")
type Base struct{
	ID  uuid.UUID `gorm:"type:varchar(36);primary_key"`
	CreateBy string
	CreateAt time.Time
	DeleteAt *time.Time
}
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	scope.SetColumn("CreateBy", "neha")
	scope.SetColumn("CreateAt", time.Now())
	return nil
}