package investmentHandler

import (
	"bankingV2/dto/DtoInvestment"
	"encoding/json"
	"log"
	"net/http"
)

func (ih InvestmentHandlerGorm) InvestmentCompanyCreate(w http.ResponseWriter, r *http.Request) {
	var companyRequest DtoInvestment.NewCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&companyRequest); err != nil {
		log.Println("Error while decoding investment risk level request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response, err := ih.S.CreateCompany(companyRequest)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}
