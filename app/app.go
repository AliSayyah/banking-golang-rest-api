package app

import (
	"github.com/AliSayyah/banking/domain"
	"github.com/AliSayyah/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	log.Println("Starting server on port 8000")

	router := mux.NewRouter()
	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define routes
	router.HandleFunc("/customers", ch.getAllCostumers).Methods(http.MethodGet)

	// Start the server
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
