package service

import (
	"banking/customErrors"
	"banking/domain"
	"banking/dto"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service banking/service CustomerService
type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *customErrors.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *customErrors.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *customErrors.AppError) {
	var response []dto.CustomerResponse

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}

	return response, nil
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
