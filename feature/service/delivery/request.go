package delivery

import (
	"RESTAPILoundry/domain"
	"time"
)

type ServicesInsertRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name"`
	Price int32  `json:"price" form:"price"`
}

func (si *ServicesInsertRequest) ToDomain() domain.Service {
	return domain.Service{
		Name:      si.Name,
		Price:     int(si.Price),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
