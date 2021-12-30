package domain

import "banking/customErrors"

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

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *customErrors.AppError)

	/**
	the following returns a customer if exists, else returns nil
	in order to to that - it returns the pointer of a Customer
	*/
	ById(string) (*Customer, *customErrors.AppError)
}
