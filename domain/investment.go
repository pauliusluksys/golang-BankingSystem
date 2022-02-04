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
	CustomerName                sql.NullString `db:"customer_name"`
	InvestmentID                uint           `db:"investment_id"`
	InvestedAmount              uint           `db:"invested_amount"`
	WithdrawnState              string         `db:"withdrawn_state"`
	CustomerInvestmentCreatedAt sql.NullTime   `db:"customer_investment_created_at"`
	CustomerInvestmentDeletedAt sql.NullTime   `db:"customer_investment_deleted_at"`
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
		ci.InvestedAmount,
		ci.WithdrawnState,
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
func (cI CustomerInvestment) ByCustomerToDto() DtoInvestment.ByCustomerResponse {
	return DtoInvestment.ByCustomerResponse{
		CustomerId:   cI.CustomerID,
		CustomerName: cI.CustomerName.String,
	}
}
func (cI CustomerInvestment) ByInvestmentToDto() DtoInvestment.ByInvestmentResponse {
	cInvCreatedAt := utils.TimeToString(cI.InvestmentCreatedAt)
	InvUpdatedAt := utils.TimeToString(cI.InvestmentUpdatedAt)
	return DtoInvestment.ByInvestmentResponse{
		InvestmentId:           cI.InvestmentID,
		InvestmentTitle:        cI.InvestmentTitle,
		InvestmentCategoryName: cI.CategoryName.String,
		InvestmentCompanyName:  cI.CompanyName.String,
		InvestmentRiskLevel:    cI.RiskLevelName.String,
		InvestmentCreatedAt:    cInvCreatedAt,
		InvestmentUpdatedAt:    InvUpdatedAt,
	}
}
func (cI CustomerInvestment) CustomersInvestmentsToDto() DtoInvestment.CustomerInvestmentResponse {
	cInvCreatedAt := utils.TimeToString(cI.CustomerInvestmentCreatedAt)
	InvCreatedAt := utils.TimeToString(cI.InvestmentCreatedAt)
	InvUpdatedAt := utils.TimeToString(cI.InvestmentUpdatedAt)
	InvDeletedAt := utils.TimeToString(cI.InvestmentDeletedAt)

	return DtoInvestment.CustomerInvestmentResponse{

		AmountInvested:              cI.InvestedAmount,
		IsWithdrawn:                 cI.WithdrawnState,
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
func OverwriteTableNameInvestments() {

	_ = InvestmentGorm.TableName
}

type Tabler interface {
	TableName() string
}

func (InvestmentGorm) TableName() string {
	return "investments"
}
