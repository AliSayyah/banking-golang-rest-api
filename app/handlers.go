package app

import (
	"encoding/json"
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

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, _ *http.Request) {
	// handler for /customers
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())

		return
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// handler for /customers/{id}
	vars := mux.Vars(r)
	var id, err = strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	customer, errs := ch.service.GetCustomer(id)
	if errs != nil {
		writeResponse(w, http.StatusNotFound, errs.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
