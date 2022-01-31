package investmentHandler

import (
	"bankingV2/dto/DtoInvestment"
	"bankingV2/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type InvestmentHandlerGorm struct {
	S service.InvestmentServiceGorm
}
type InvestmentHandler struct {
	S service.InvestmentService
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
func (ih InvestmentHandler) GetAllCustomerInvestments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	response, err := ih.S.GetAllCustomerInvestments(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		writeResponse(w, http.StatusOK, response)
	}
}
func (ih InvestmentHandler) CustomerInvestmentCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var ciRequest DtoInvestment.NewCustomerInvestmentRequest
	err := json.NewDecoder(r.Body).Decode(&ciRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	}
	custId64, err := strconv.ParseUint(customerId, 10, 64)
	ciRequest.CustomerID = uint(custId64)

	if err != nil {
		log.Println("Error while decoding new customer investment request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response, err := ih.S.CreateCustomerInvestment(ciRequest)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}

	}
}
func (ih InvestmentHandlerGorm) InvestmentsCreate(w http.ResponseWriter, r *http.Request) {
	var request DtoInvestment.NewInvestmentRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error while decoding new investment request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response, err := ih.S.CreateInvestment(request)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}

	}
}
