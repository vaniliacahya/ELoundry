package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Service struct {
	ID        int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceUseCase interface {
	AddService(newService Service) (Service, error)
	DeleteService(id int) bool
	UpdateService(id int, updateProfile Service) (Service, error)
	GetAllS() ([]Service, error)
	GetSpecificServices(serviceID int) ([]Service, error)
}

type ServiceData interface {
	Insert(newService Service) Service
	Delete(serviceID int) bool
	Update(serviceID int, updatedData Service) Service
	GetAll() []Service
	GetServiceID(serviceID int) []Service
}

type ServiceHandler interface {
	InsertServ() echo.HandlerFunc
	GetAllServ() echo.HandlerFunc
	UpdateServ() echo.HandlerFunc
	DeleteServ() echo.HandlerFunc
	GetServID() echo.HandlerFunc
}
