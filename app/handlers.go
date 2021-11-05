package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, _ *http.Request) {
	// handler for /greet
	_, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		return
	}
}

func getAllCostumers(w http.ResponseWriter, r *http.Request) {
	// handler for /customers
	customers := []Customer{
		{
			Name:    "John",
			City:    "New York",
			ZipCode: "10001",
		},
		{
			Name:    "Jane",
			City:    "New York",
			ZipCode: "10001",
		},
		{
			Name:    "Joe",
			City:    "New York",
			ZipCode: "10001",
		},
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
