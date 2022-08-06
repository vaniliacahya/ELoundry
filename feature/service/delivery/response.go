package delivery

import (
	"RESTAPILoundry/domain"
)

type ServicesResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name"`
	Price int32  `json:"price" form:"price"`
}

func FromDomain(data domain.Service) ServicesResponse {
	var res ServicesResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.Price = int32(data.Price)
	return res
}
