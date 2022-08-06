package data

import (
	"RESTAPILoundry/domain"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name  string `json:"name" form:"name" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
}

func (s *Service) ToDomain() domain.Service {
	return domain.Service{
		ID:        int(s.ID),
		Name:      s.Name,
		Price:     s.Price,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func ParseToArr(arr []Service) []domain.Service {
	var res []domain.Service

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Service) Service {
	var res Service
	res.Name = data.Name
	res.Price = data.Price
	return res
}
