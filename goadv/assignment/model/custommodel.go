package model

import (uuid"github.com/satori/go.uuid")
type CustomModel struct{
	ID  uuid.UUID `gorm:"type:varchar(36);primary_key"`
	CreateBy string
}