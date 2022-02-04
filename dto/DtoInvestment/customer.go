package DtoInvestment

type NewCustomerInvestmentResponse struct {
	CustomerID      uint `json:"customer_id"`
	InvestmentID    uint `json:"investment_id"`
	AmountInvested  uint
	IsWithdrawn     string
	CreatedAt       string
	CustomerName    string
	InvestmentTitle string
}
type AllCustomersResponse struct {
	Customers []AllCustomerInvestmentResponse
}
type AllCustomerInvestmentResponse struct {
	UniqueInvestmentsCount     int `json:"unique_investments_count"`
	CustomerInvestmentResponse []CustomerInvestmentResponse
}
type CustomerInvestmentResponse struct {
	AmountInvested              uint   `json:"amount_invested"`
	IsWithdrawn                 string `json:"is_withdrawn"`
	CustomerInvestmentCreatedAt string `json:"customer_investment_created_at,omitempty"`
	InvestmentID                uint   `json:"investment_id"`
	InvestmentCreatedAt         string `json:"created_at"`
	InvestmentUpdatedAt         string `json:"updated_at"`
	InvestmentDeletedAt         string `json:"deleted_at,omitempty"`
	InvestmentTitle             string `json:"investment_title"`
	CategoryInvestmentID        int64  `json:"category_investment_id"`
	CompanyInvestmentID         int64  `json:"company_investment_id"`
	RiskLevelInvestmentID       int64  `json:"risk_level_investment_id"`
	CompanyName                 string `json:"company_name"`
	CategoryName                string `json:"category_name"`
	RiskLevelName               string `json:"risk_level_name"`
}
type ByInvestmentResponse struct {
	InvestmentId                  uint   `json:"investment_id"`
	InvestmentTitle               string `json:"investment_title"`
	InvestmentCreatedAt           string `json:"created_at"`
	InvestmentUpdatedAt           string `json:"updated_at"`
	InvestmentCompanyName         string `json:"company_name"`
	InvestmentCategoryName        string `json:"category_name"`
	InvestmentRiskLevel           string `json:"risk_level_name"`
	CustomerInvestmentLastCreated string `json:"last_made_investment_at"`
	TotalInvestments              int    `json:"total_investments"`
	FullAmountInvested            int    `json:"full_amount_invested"`
	CustomerInvestmentResponse    []CustomerInvestmentResponse
}
type ByCustomerResponse struct {
	CustomerId           uint   `json:"customer_id"`
	CustomerName         string `json:"customer_name"`
	ByInvestmentResponse []ByInvestmentResponse
}

//type OneCustomerInvestmentResponse struct {
//}
