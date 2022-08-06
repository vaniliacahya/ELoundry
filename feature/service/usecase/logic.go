package usecase

import (
	"RESTAPILoundry/domain"
	"errors"

	validator "github.com/go-playground/validator/v10"
)

type serviceUseCase struct {
	serviceData domain.ServiceData
	validate    *validator.Validate
}

func New(ud domain.ServiceData, v *validator.Validate) domain.ServiceUseCase {
	return &serviceUseCase{
		serviceData: ud,
		validate:    v,
	}
}

func (sd *serviceUseCase) AddService(newService domain.Service) (domain.Service, error) {

	res := sd.serviceData.Insert(newService)

	if res.ID == 0 {
		return domain.Service{}, errors.New("error insert data")
	}
	return res, nil
}

func (sd *serviceUseCase) GetSpecificServices(serviceID int) ([]domain.Service, error) {
	res := sd.serviceData.GetServiceID(serviceID)
	if serviceID == -1 {
		return nil, errors.New("error get data")
	}

	return res, nil
}

func (sd *serviceUseCase) GetAllS() ([]domain.Service, error) {
	res := sd.serviceData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (sd *serviceUseCase) UpdateService(serviceID int, updateData domain.Service) (domain.Service, error) {

	if serviceID == -1 {
		return domain.Service{}, errors.New("invalid data")
	}
	result := sd.serviceData.Update(serviceID, updateData)

	if result.ID == 0 {
		return domain.Service{}, errors.New("error update data")
	}
	return result, nil
}

func (sd *serviceUseCase) DeleteService(serviceID int) bool {
	res := sd.serviceData.Delete(serviceID)

	if !res {
		return false
	}

	return true
}
