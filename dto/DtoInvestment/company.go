package DtoInvestment

type CompanyRequest struct {
	Name string `json:"name"`
}
type NewCompanyResponse struct {
	ID   uint   `json:"company_id"`
	Name string `json:"name"`
}
type NewCompanyRequest struct {
	Name string `json:"name"`
}
