package domain

import (
	"banking/customErrors"
	"banking/dto"
)

/**
this is the relation to the database table
*/
type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string

	/**
	the database table contains customer_id but our struct contains id
	so the code `db:"customer_id"`
	makes the relation between db and code
	*/
}

func (c Customer) statusAsText() string {
	// DB Value: 0 --> status inactive
	// DB Value: 1 --> status active

	statusAsText := "active"

	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *customErrors.AppError)

	/**
	the following returns a customer if exists, else returns nil
	in order to to that - it returns the pointer of a Customer
	*/
	ById(string) (*Customer, *customErrors.AppError)
}
