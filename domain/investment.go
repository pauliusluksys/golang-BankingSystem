package domain

import (
	"bankingV2/dto/DtoInvestment"
	"bankingV2/utils"
	"database/sql"
	"gorm.io/gorm"
)

type InvestmentGorm struct {
	gorm.Model
	Title                 string      `gorm:"uniqueIndex:idx_name,title"`
	CategoryInvestmentID  uint        `gorm:"uniqueIndex:idx_name,category_investment_id"`
	CompanyInvestmentID   uint        `gorm:"uniqueIndex:idx_name,company_investment_id"`
	RiskLevelInvestmentID uint        `gorm:"uniqueIndex:idx_name,risk_level_investment_id"`
	Customers             []*Customer `gorm:"many2many:customer_investments"`
}
type Investment struct {
	ID                    uint         `db:"id"`
	Title                 string       `db:"title"`
	CreatedAt             sql.NullTime `db:"created_at"`
	UpdatedAt             sql.NullTime `db:"updated_at"`
	DeletedAt             sql.NullTime `db:"deleted_at"`
	CategoryInvestmentID  uint         `db:"category_investment_id"`
	CompanyInvestmentID   uint         `db:"company_investment_id"`
	RiskLevelInvestmentID uint         `db:"risk_level_investment_id"`
}

type CategoryInvestment struct {
	gorm.Model
	Name           string `gorm:"unique"`
	InvestmentGorm InvestmentGorm
}
type CompanyInvestment struct {
	gorm.Model
	Name           string `gorm:"unique"`
	InvestmentGorm InvestmentGorm
}
type RiskLevelInvestment struct {
	gorm.Model
	Name           string `gorm:"unique"`
	InvestmentGorm InvestmentGorm
}
type CustomerInvestment struct {
	CustomerID                  uint           `db:"customer_id"`
	InvestmentID                uint           `db:"investment_id"`
	AmountInvested              uint           `db:"amount_invested"`
	IsWithdrawn                 bool           `db:"is_withdrawn"`
	CustomerInvestmentCreatedAt sql.NullTime   `db:"customer_investment_created_at"`
	CustomerInvestmentDeletedAt sql.NullTime   `db:"customer_investment_deleted_at"`
	CustomerName                sql.NullString `db:"customer_name"`
	InvestmentTitle             string         `db:"investment_title"`
	CategoryInvestmentID        sql.NullInt64  `db:"category_investment_id"`
	CompanyInvestmentID         sql.NullInt64  `db:"company_investment_id"`
	RiskLevelInvestmentID       sql.NullInt64  `db:"risk_level_investment_id"`
	CompanyName                 sql.NullString `db:"company_name"`
	CategoryName                sql.NullString `db:"category_name"`
	RiskLevelName               sql.NullString `db:"risk_level_name"`
	InvestmentCreatedAt         sql.NullTime   `db:"investment_created_at"`
	InvestmentUpdatedAt         sql.NullTime   `db:"investment_updated_at"`
	InvestmentDeletedAt         sql.NullTime   `db:"investment_deleted_at"`
}
type CustomerInvestmentsCount struct {
	Total int
}

func (ci CustomerInvestment) NewCustomerInvestmentToDto() DtoInvestment.NewCustomerInvestmentResponse {
	return DtoInvestment.NewCustomerInvestmentResponse{
		ci.InvestmentID,
		ci.InvestmentID,
		ci.AmountInvested,
		ci.IsWithdrawn,
		ci.CustomerInvestmentCreatedAt.Time.String(),
		ci.CustomerName.String,
		ci.InvestmentTitle,
	}
}
func (investment InvestmentGorm) NewInvestmentToDto() DtoInvestment.NewInvestmentResponse {
	return DtoInvestment.NewInvestmentResponse{
		ID: investment.ID,
	}
}
func (category CategoryInvestment) NewCategoryToDto() DtoInvestment.NewCategoryResponse {
	return DtoInvestment.NewCategoryResponse{
		ID: category.ID,
	}
}
func (riskLevel RiskLevelInvestment) NewRiskLevelToDto() DtoInvestment.NewRiskLevelResponse {
	return DtoInvestment.NewRiskLevelResponse{
		ID: riskLevel.ID,
	}
}
func (company CompanyInvestment) NewCompanyToDto() DtoInvestment.NewCompanyResponse {
	return DtoInvestment.NewCompanyResponse{
		ID: company.ID,
	}
}
func (cI CustomerInvestment) CustomerInvestmentsToDto() DtoInvestment.CustomerInvestmentResponse {
	cInvCreatedAt := utils.TimeToString(cI.CustomerInvestmentCreatedAt)
	InvCreatedAt := utils.TimeToString(cI.InvestmentCreatedAt)
	InvUpdatedAt := utils.TimeToString(cI.InvestmentUpdatedAt)
	InvDeletedAt := utils.TimeToString(cI.InvestmentDeletedAt)

	return DtoInvestment.CustomerInvestmentResponse{
		AmountInvested:              cI.AmountInvested,
		IsWithdrawn:                 cI.IsWithdrawn,
		CustomerInvestmentCreatedAt: cInvCreatedAt,
		InvestmentID:                cI.InvestmentID,
		InvestmentCreatedAt:         InvCreatedAt,
		InvestmentUpdatedAt:         InvUpdatedAt,
		InvestmentDeletedAt:         InvDeletedAt,
		InvestmentTitle:             cI.InvestmentTitle,
		CategoryInvestmentID:        cI.CategoryInvestmentID.Int64,
		CompanyInvestmentID:         cI.CompanyInvestmentID.Int64,
		RiskLevelInvestmentID:       cI.RiskLevelInvestmentID.Int64,
		CompanyName:                 cI.CompanyName.String,
		CategoryName:                cI.CategoryName.String,
		RiskLevelName:               cI.RiskLevelName.String,

		//AmountInvested:              cI.AmountInvested,
		//IsWithdrawn:                 cI.IsWithdrawn,
		//CustomerInvestmentCreatedAt: cI.CustomerInvestmentCreatedAt.Time,
		//InvestmentID:                cI.InvestmentID,
		//InvestmentCreatedAt:         *cI.InvestmentCreatedAt,
		//InvestmentUpdatedAt:         *cI.InvestmentUpdatedAt,
		//InvestmentDeletedAt:         cI.InvestmentUpdatedAt,
		//InvestmentTitle:             cI.InvestmentTitle,
		//CategoryInvestmentID:        cI.CategoryInvestmentID.Int64,
		//CompanyInvestmentID:         cI.CompanyInvestmentID.Int64,
		//RiskLevelInvestmentID:       cI.RiskLevelInvestmentID.Int64,
		//CompanyName:                 cI.CompanyName.String,
		//CategoryName:                cI.CategoryName.String,
		//RiskLevelName:               cI.RiskLevelName.String,
	}
}
