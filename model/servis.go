package model

import (
	"gorm.io/gorm"
)

type Servis struct {
	ID          int    `json:"id" form:"id" gorm:"prmaryKey;autoIncrement"`
	ServiceName string `json:"nama" form:"nama"`
}

type ServisModel struct {
	DB *gorm.DB
}
