package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         int    `json:"id" form:"id" gorm:"prmaryKey;autoIncrement"`
	Nama       string `json:"nama" form:"nama"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	UserAddres string `json:"address" form:"address"`
}

type UserModel struct {
	DB *gorm.DB
}
