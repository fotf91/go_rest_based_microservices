package service

import (
	"banking/customErrors"
	"banking/domain"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *customErrors.AppError)
	GetCustomer(string) (*domain.Customer, *customErrors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *customErrors.AppError) {
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *customErrors.AppError) {
	return s.repo.ById(id)
}

/**
function that instantiates the DefaultCustomerService
*/
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
