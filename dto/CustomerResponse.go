package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipconde"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `status:"status"`
}
