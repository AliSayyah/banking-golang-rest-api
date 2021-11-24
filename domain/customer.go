package domain

import "github.com/AliSayyah/banking/errs"

type Customer struct {
	ID          int `db:"customer_id"`
	Name        string
	City        string
	Status      string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
}

type CustomerRepository interface {
	// FindAll status == 1 means active
	// status == 0 means inactive
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(int) (*Customer, *errs.AppError)
}
