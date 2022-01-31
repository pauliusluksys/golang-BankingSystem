package domain

import (
	"bankingV2/dto"
	"bankingV2/errs"
)

type Customer struct {
	ID          uint   `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}
type CustomerGorm struct {
	ID          uint
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
	Investments []InvestmentGorm `gorm:"many2many:customer_investments"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}
func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(UrlStatus string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}
