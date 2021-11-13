package service

import (
	"github.com/AliSayyah/banking/domain"
	"github.com/AliSayyah/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(id int) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id int) (*domain.Customer, *errs.AppError) {
	return s.repo.FindByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
