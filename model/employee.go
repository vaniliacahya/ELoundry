package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID              string `json:"id" form:"id" gorm:"prmaryKey;"`
	EmployeeName    string `json:"nama" form:"nama"`
	EmployeeAddress string `json:"address" form:"address"`
}

type EmployeeModel struct {
	DB *gorm.DB
}
