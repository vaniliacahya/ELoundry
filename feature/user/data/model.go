package data

import (
	"RESTAPILoundry/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"default:users"`
	Photo    string `json:"image_url"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		UserName:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FullName:  u.FullName,
		Role:      u.Role,
		Photo:     u.Photo,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Username = data.UserName
	res.Email = data.Email
	res.Password = data.Password
	res.FullName = data.FullName
	res.Photo = data.Photo
	res.Role = data.Role
	return res
}
