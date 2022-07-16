package model

import (
	"gorm.io/gorm"
)

type Servis struct {
	gorm.Model
	IDservice  int     `json:"id" form:"id" gorm:"primaryKey;"`
	Qantity    float64 `json:"berat" form:"berat"`
	IDKategory int
}

type ServisModel struct {
	DB *gorm.DB
}
