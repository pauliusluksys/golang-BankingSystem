package investmentHandler

import (
	"bankingV2/dto/DtoInvestment"
	"encoding/json"
	"log"
	"net/http"
)

func (ih InvestmentHandlerGorm) InvestmentRiskLevelCreate(w http.ResponseWriter, r *http.Request) {
	var riskLevelRequest DtoInvestment.NewRiskLevelRequest
	if err := json.NewDecoder(r.Body).Decode(&riskLevelRequest); err != nil {
		log.Println("Error while decoding investment risk level request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response, err := ih.S.CreateRiskLevel(riskLevelRequest)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}
