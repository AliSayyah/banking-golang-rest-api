package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/AliSayyah/banking/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCostumers(w http.ResponseWriter, r *http.Request) {
	// handler for /customers
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	}
}
