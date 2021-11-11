package domain

type Customer struct {
	ID          int
	Name        string
	City        string
	Status      string
	Zipcode     string
	DateOfBirth string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
