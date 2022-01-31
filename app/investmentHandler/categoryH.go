package investmentHandler

import (
	"bankingV2/dto/DtoInvestment"
	"encoding/json"
	"log"
	"net/http"
)

func (ih InvestmentHandlerGorm) InvestmentCategoryCreate(w http.ResponseWriter, r *http.Request) {
	var categoryRequest DtoInvestment.NewCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		log.Println("Error while decoding investment risk level request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response, err := ih.S.CreateCategory(categoryRequest)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}
