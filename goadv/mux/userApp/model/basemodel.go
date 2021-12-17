package model

import (uuid"github.com/satori/go.uuid"
"time")
type Base struct{
	ID  uuid.UUID `gorm:"type:varchar(36);primary_key"`
	CreateBy string
	CreateAt time.Time
	DeleteAt *time.Time
}