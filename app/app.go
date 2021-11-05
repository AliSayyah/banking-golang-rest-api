package app

import (
	"log"
	"net/http"
)

func Start() {
	log.Println("Starting server on port 8000")

	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCostumers)

	// Start the server
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
