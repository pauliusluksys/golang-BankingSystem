package app

import (
	"bankingV2/dto"
	"bankingV2/service"
	"bankingV2/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type JobHandler struct {
	service service.JobService
}
type JobHandlerGorm struct {
	service service.JobServiceGorm
}

func (jh JobHandlerGorm) DeleteJob(w http.ResponseWriter, r *http.Request) {
	var request dto.JobRequestGorm
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		jobDeleted, appError := jh.service.DeleteJob(request)
		if err != nil {
			fmt.Println(err)
		}
		if appError != nil {
			log.Println("error")
			writeResponse(w, appError.Code, appError.Message)
		} else {
			//log.Println("continue on")
			writeResponse(w, http.StatusCreated, jobDeleted)
		}
	}
}
func (jh JobHandlerGorm) UpdateJob(w http.ResponseWriter, r *http.Request) {
	var request dto.JobRequestGorm
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		job, appError := jh.service.UpdateJob(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			//log.Println("continue on")
			writeResponse(w, http.StatusCreated, job)
		}
	}
}
func (jh JobHandlerGorm) NewJob(w http.ResponseWriter, r *http.Request) {
	var request dto.JobRequestGorm
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		job, appError := jh.service.NewJob(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			//log.Println("continue on")
			writeResponse(w, http.StatusCreated, job)
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
