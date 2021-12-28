package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zip_code"`

// 	/**
// 	code: Name    string `json:"full_name"`
// 	means that the name of the attribute is Name
// 	but when converted to json is shown as full_name
// 	----
// 	code: Name    string `json:"full_name" xml:"name"`
// 	same as above for xml
// 	*/
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/**
	  w http.ResponseWriter - sends the response back to the client
	  r *http.Request - request coming to the server
	*/

	customers, _ := ch.service.GetAllCustomer()

	// check if the request headers want to retrieve xml or json formatted response
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
