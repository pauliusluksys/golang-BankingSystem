package domain

import (
	"bankingV2/dto"
	"bankingV2/errs"
	"gorm.io/gorm"
	"log"
)

type InvestmentRepositoryDbGorm struct {
	Client *gorm.DB
}

func (d InvestmentRepositoryDbGorm) FindAll() ([]InvestmentGorm, *errs.AppError) {
	var investment []InvestmentGorm
	result := d.Client.Find(&investment)
	if result.Error != nil {
		log.Println("Error with database:", result.Error)
		return nil, errs.GormQueryError("Unexpected database error")
	}
	return investment, nil
}
func NewInvestmentRepositoryDbGorm(dbClient *gorm.DB) InvestmentRepositoryDbGorm {
	return InvestmentRepositoryDbGorm{dbClient}
}

func (Investment InvestmentGorm) ToInvestmentResponseGormDto() dto.InvestmentResponseGorm {

	return dto.InvestmentResponseGorm{
		ID:    Investment.ID,
		Title: Investment.Title,
	}
}
