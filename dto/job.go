package dto

import (
	"bankingV2/errs"
	"strings"
)

type JobRequest struct {
	Id string `json:"job_id"`
}
type AllJobsRequest struct {
	Location     string `json:"location"`
	CategoryName string `json:"category_name"`
}
type JobResponse struct {
	Id           string `json:"job_id,omitempty"`
	Title        string `json:"job_title,omitempty"`
	Description  string `json:"job_description,omitempty"`
	CategoryName string `json:"job_category,omitempty"`
	ApplyUntil   string `json:"apply_until,omitempty"`
	City         string `json:"city,omitempty"`
	IsExpired    bool   `json:"is_expired"`
}

func validateAndReturnFilterMap(filter string) (map[string]string, *errs.AppError) {
	splits := strings.Split(filter, ".")
	//if len(splits) != 2 {
	//	return nil, errors.New("malformed filter as URL param")
	//}
	field, value := splits[0], splits[1]
	filterName := []string{"category_name", "location"}
	if !stringInSlice(filterName, field) {
		//return nil, errors.New("unknown field in filter query parameter")
	}
	return map[string]string{field: value}, nil
}

func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}
	return false
}
