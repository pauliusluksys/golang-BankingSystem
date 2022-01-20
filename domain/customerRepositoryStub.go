package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"12", "Me", "Vilno", "54543", "2021-12-11", "1"},
		{"13", "mom", "Vilno", "54543", "2021-12-11", "1"},
		{"14", "dad", "Vilno", "54543", "2021-12-11", "1"},
		{"15", "son", "Vilno", "54543", "2021-12-11", "1"},
	}
	return CustomerRepositoryStub{customers}
}
