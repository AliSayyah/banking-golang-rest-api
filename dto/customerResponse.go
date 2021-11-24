package dto

type CustomerResponse struct {
	ID          int    `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Status      string `json:"status"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth"`
}
