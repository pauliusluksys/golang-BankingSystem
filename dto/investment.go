package dto

type InvestmentResponseGorm struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	Company   string `json:"company"`
	RiskLevel string `json:"risk_level"`
}
