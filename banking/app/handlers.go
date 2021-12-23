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

	/**
	code: Name    string `json:"full_name"`
	means that the name of the attribute is Name
	but when converted to json is shown as full_name
	----
	code: Name    string `json:"full_name" xml:"name"`
	same as above for xml
	*/
}

func greet(w http.ResponseWriter, r *http.Request) {
	/**
	  w http.ResponseWriter - sends the response back to the client
	  r *http.Request - request coming to the server
	*/
	fmt.Fprint(w, "Hello World!!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/**
	  w http.ResponseWriter - sends the response back to the client
	  r *http.Request - request coming to the server
	*/
	customers := []Customer{
		{Name: "Fotis", City: "Athens", ZipCode: "12345"},
		{Name: "Alex", City: "Lp", ZipCode: "12345"},
	}

	// check if the request headers want to retrieve xml or json formatted response
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
