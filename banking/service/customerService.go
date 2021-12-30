package service

import (
	"banking/customErrors"
	"banking/domain"
	"banking/dto"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *customErrors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *customErrors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *customErrors.AppError) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *customErrors.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

/**
function that instantiates the DefaultCustomerService
*/
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
