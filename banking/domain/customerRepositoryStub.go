package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Fotis", City: "Athens", Zipcode: "1245", DateofBirth: "1991-07-19", Status: "1"},
		{Id: "2", Name: "Alex", City: "Athens", Zipcode: "1245", DateofBirth: "1981-10-19", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
