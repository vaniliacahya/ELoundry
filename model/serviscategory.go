package model

import (
	"gorm.io/gorm"
)

type ServisCategory struct {
	ID           int    `json:"id" form:"id" gorm:"prmaryKey;autoIncrement"`
	CategoryName string `json:"nama" form:"nama"`
}

type ServisCatModel struct {
	DB *gorm.DB
}
