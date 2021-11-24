package service

import (
	"github.com/AliSayyah/banking/domain"
	"github.com/AliSayyah/banking/dto"
	"github.com/AliSayyah/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	switch {
	case status == "active":
		status = "1"
	case status == "inactive":
		status = "0"
	default:
		status = ""
	}
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var customers []dto.CustomerResponse
	for _, c := range c {
		customers = append(customers, c.ToDTO())
	}
	return customers, nil
}
func (s DefaultCustomerService) GetCustomer(id int) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDTO()
	return &response, nil

}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
