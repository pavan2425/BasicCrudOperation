package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "101", Name: "pavan", Dateofbirth: "24/06/1995", City: "bangalore", Zipcode: "570000", Status: "1"},
		{Id: "102", Name: "sulu", Dateofbirth: "24/06/1965", City: "bangalore", Zipcode: "570000", Status: "1"},

		{Id: "103", Name: "dukku", Dateofbirth: "24/06/1975", City: "bangalore", Zipcode: "570000", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
