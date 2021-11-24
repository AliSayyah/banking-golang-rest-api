package app

import (
	"github.com/AliSayyah/banking/domain"
	"github.com/AliSayyah/banking/logger"
	"github.com/AliSayyah/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	logger.Info("Starting server on port 8000")

	router := mux.NewRouter()
	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// Start the server
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
