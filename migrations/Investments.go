package migrations

import (
	"bankingV2/domain"
	"fmt"
	"gorm.io/gorm"
)

func MigrateInvestments(db *gorm.DB) {
	err := db.AutoMigrate(&domain.InvestmentCategory{}, &domain.InvestmentCompany{}, &domain.RiskLevel{}, &domain.InvestmentGorm{}, &domain.CustomerInvestment{})
	if err != nil {
		fmt.Println("something went wrong with Gorm migration:   ", err)
	}
}
