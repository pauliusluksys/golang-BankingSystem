package DtoInvestment

import (
	"bankingV2/errs"
)

type NewInvestmentRequest struct {
	Title                string `json:"title"`
	InvestmentCategoryID uint   `json:"category_id"`
	InvestmentCompanyID  uint   `json:"company_id"`
	RiskLevelID          uint   `json:"risk_level_id"`
	CustomerID           uint   `json:"customer_id"`
}
type NewCustomerInvestmentRequest struct {
	CustomerID     uint
	InvestmentID   uint `json:"investment_id"`
	AmountInvested uint `json:"amount_invested"`
	IsWithdrawn    bool `json:"is_withdrawn"`
}

type NewInvestmentResponse struct {
	ID uint
}

func (r NewInvestmentRequest) Validate() *errs.AppError {
	return nil
}
