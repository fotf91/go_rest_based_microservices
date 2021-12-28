package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/api/time", getTime).Queries("tz", "{tz}").Methods(http.MethodGet) // route with query parameter
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)                       // same route without query parameter, because it is optional

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
