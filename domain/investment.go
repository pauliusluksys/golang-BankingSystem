package domain

import (
	"gorm.io/gorm"
	"time"
)

type InvestmentGorm struct {
	gorm.Model
	Title                string
	InvestmentCategoryID uint
	InvestmentCompanyID  uint
	RiskLevelID          uint
	Customers            []*Customer `gorm:"many2many:customer_investments"`
}
type InvestmentCategory struct {
	gorm.Model
	Name           string
	InvestmentGorm InvestmentGorm
}
type InvestmentCompany struct {
	gorm.Model
	Name           string
	InvestmentGorm InvestmentGorm
}
type RiskLevel struct {
	gorm.Model
	Name           string
	InvestmentGorm InvestmentGorm
}
type CustomerInvestment struct {
	CustomerID     int `gorm:"primaryKey"`
	InvestmentID   int `gorm:"primaryKey"`
	AmountInvested int
	IsWithdrawn    string
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
