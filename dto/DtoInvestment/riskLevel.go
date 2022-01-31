package DtoInvestment

//type RiskLevelRequest struct {
//	Name string	`json:"name"`
//}
type NewRiskLevelRequest struct {
	Name string `json:"name"`
}
type NewRiskLevelResponse struct {
	ID uint `json:"id"`
}
