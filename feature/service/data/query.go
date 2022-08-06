package data

import (
	"RESTAPILoundry/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type serviceData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ServiceData {
	return &serviceData{
		db: db,
	}
}

func (sd *serviceData) Insert(newService domain.Service) domain.Service {
	cnv := ToLocal(newService)
	err := sd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Service{}
	}
	return cnv.ToDomain()
}

func (sd *serviceData) Update(serviceID int, updatedService domain.Service) domain.Service {
	cnv := ToLocal(updatedService)
	err := sd.db.Model(cnv).Where("ID = ?", serviceID).Updates(updatedService)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Service{}
	}
	cnv.ID = uint(serviceID)
	return cnv.ToDomain()
}

func (sd *serviceData) Delete(serviceID int) bool {
	err := sd.db.Where("ID = ?", serviceID).Delete(&Service{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (sd *serviceData) GetAll() []domain.Service {
	var data []Service
	err := sd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (sd *serviceData) GetServiceID(serviceID int) []domain.Service {
	var data []Service
	err := sd.db.Where("ID = ?", serviceID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
