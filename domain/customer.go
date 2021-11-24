package domain

import (
	"github.com/AliSayyah/banking/dto"
	"github.com/AliSayyah/banking/errs"
)

type Customer struct {
	ID          int `db:"customer_id"`
	Name        string
	City        string
	Status      string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Status:      c.statusAsText(),
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
	}
}

type CustomerRepository interface {
	// FindAll status == 1 means active
	// status == 0 means inactive
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(int) (*Customer, *errs.AppError)
}
