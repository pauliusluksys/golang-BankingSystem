package app

import (
	"bankingV2/dto"
	"bankingV2/service"
	"bankingV2/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type JobHandler struct {
	service service.JobService
}

func (jh JobHandler) NewJob(w http.ResponseWriter, r *http.Request) {
	var request dto.JobRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := jh.service.NewJob(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			//log.Println("continue on")
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
func (jh JobHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["job_id"]

	job, err := jh.service.GetById(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, job)
	}
}
func (jh JobHandler) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	var request dto.AllJobsRequest
	filterMap := map[string]string{}
	filterMap["location"] = r.URL.Query().Get("location")
	filterMap["category"] = r.URL.Query().Get("team")
	filterMap = utils.CleanMap(filterMap)
	fmt.Println(filterMap)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		fmt.Println(request)
		jobs, err := jh.service.GetAllJobs(request, filterMap)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, jobs)
		}
	}
}
