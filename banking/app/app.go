package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Queries("status", "{status}").Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
