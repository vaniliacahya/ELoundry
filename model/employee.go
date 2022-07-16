package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	ID             int    `json:"id" form:"id" gorm:"prmaryKey;autoIncrement"`
	EmployeeName   string `json:"nama" form:"nama"`
	EmployeeAddres string `json:"addres" form:"addres"`
}

type EmployeeModel struct {
	DB *gorm.DB
}
