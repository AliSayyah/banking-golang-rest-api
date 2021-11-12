package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func _() CustomerRepositoryStub {
	return CustomerRepositoryStub{
		customers: []Customer{
			{
				ID:          1,
				Name:        "John",
				City:        "New York",
				Status:      "Active",
				ZipCode:     "10001",
				DateOfBirth: "01/01/1990",
			},
			{
				ID:          2,
				Name:        "Jane",
				City:        "New York",
				Status:      "Active",
				ZipCode:     "10001",
				DateOfBirth: "01/01/1990",
			},
			{
				ID:          3,
				Name:        "Jack",
				City:        "New York",
				Status:      "Active",
				ZipCode:     "10001",
				DateOfBirth: "01/01/1990",
			},
		},
	}
}
