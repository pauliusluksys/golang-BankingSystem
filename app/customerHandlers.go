package app

import (
	"bankingV2/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("it works")
}
func getOneCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	UrlData := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(UrlData)
	if r.Header.Get("Content-Type") == "application/json" {
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

	//if r.Header.Get("Content-Type") == "application/json" {
	//	w.Header().Add("Content-Type", "application/json")
	//	err := json.NewEncoder(w).Encode(customer)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}
	//if r.Header.Get("Content-Type") == "application/xml" {
	//	w.Header().Add("Content-Type", "application/xml")
	//	xml.NewEncoder(w).Encode(customer)
	//}
}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

type TimeByZone struct {
	CurrentTime time.Time `json:"current_time"`
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	timeZoneMap := make(map[string]string)
	var timeForMap time.Time
	tzUrlVariable := r.URL.Query().Get("tz")

	if tzUrlVariable == "" {
		timeOfZone, _ := time.LoadLocation("UTC")
		timeForMap = time.Now().In(timeOfZone)
		s := fmt.Sprint(timeForMap)
		timeZoneMap["current_time"] = s

	} else {
		allStrings := strings.Split(tzUrlVariable, ",")
		for _, value := range allStrings {
			timeOfZone, err := time.LoadLocation(value)
			if err != nil {
				http.Error(w, "invalid timezone", 404)
				return
			}
			timeForMap = time.Now().In(timeOfZone)
			s := fmt.Sprint(timeForMap)
			timeZoneMap[value] = s
		}
	}
	json.NewEncoder(w).Encode(timeZoneMap)

}
