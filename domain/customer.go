package domain

type Customer struct {
	ID          int
	Name        string
	City        string
	Status      string
	ZipCode     string
	DateOfBirth string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindByID(int) (*Customer, error)
}
