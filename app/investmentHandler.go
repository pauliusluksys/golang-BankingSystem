package app

import (
	"bankingV2/service"
	"net/http"
)

type InvestmentHandlerGorm struct {
	service service.DefaultInvestmentServiceGorm
}

func (ih *InvestmentHandlerGorm) GetAllInvestments(w http.ResponseWriter, r *http.Request) {
	investments, err := ih.service.GetAllInvestments()
	if r.Header.Get("Content-Type") == "application/json" {
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, investments)
		}
	}
}
