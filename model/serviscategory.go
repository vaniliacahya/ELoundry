package model

import (
	"gorm.io/gorm"
)

type ServisCategory struct {
	gorm.Model
	IDKategory   int    `json:"id" form:"id" gorm:"primaryKey;"`
	CategoryName string `json:"nama" form:"nama"`
	PricePerKilo string `json:"price" form:"price"` //Biaya Loundry/kilo
}

type ServisCatModel struct {
	DB *gorm.DB
}
