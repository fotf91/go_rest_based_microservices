package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)

	/**
	the following returns a customer if exists, else returns nil
	in order to to that - it returns the pointer of a Customer
	*/
	ById(string) (*Customer, error)
}
