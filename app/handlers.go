package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/AliSayyah/banking/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Customer struct {
	Name    string
	City    string
	ZipCode string
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
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

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// handler for /customers/{id}
	vars := mux.Vars(r)
	var id, err = strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		return
	}
}
