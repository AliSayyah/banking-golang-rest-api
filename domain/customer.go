package domain

import "github.com/AliSayyah/banking/errs"

type Customer struct {
	ID          int
	Name        string
	City        string
	Status      string
	ZipCode     string
	DateOfBirth string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindByID(int) (*Customer, *errs.AppError)
}
