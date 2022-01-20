package domain

import (
	"bankingV2/dto"
	"bankingV2/errs"
	"fmt"
	"strings"
	"time"
)

type JobRepository interface {
	FindAll() ([]Job, *errs.AppError)
	ById(id string) (*Job, *errs.AppError)
}
type Job struct {
	Id           string `db:"id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	ApplyUntil   string `db:"apply_until"`
	CategoryName string `db:"category_title"`
	City         string `db:"location_name"`
	IsExpired    bool
}

func (job Job) ToDto() dto.JobResponse {
	today := time.Now()
	layout := "2006-01-02 15:04:05"
	str := job.ApplyUntil
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("we got a problem Heuston: ", err)
	}
	if t.After(today) {
		job.IsExpired = false
	} else {
		job.IsExpired = true
	}

	return dto.JobResponse{
		Id:           job.Id,
		Title:        job.Title,
		Description:  job.Description,
		ApplyUntil:   job.ApplyUntil,
		CategoryName: job.CategoryName,
		City:         job.City,
		IsExpired:    job.IsExpired,
	}
}

//type Filter func(job Job, filter string) bool
//
//func FilterByLocation(job Job, filter string) bool {
//	return false
//}
//func FilterByCategory(job Job, filter string) bool {
//	return false
//}
func ReturnFilterMap(filter string) map[string]string {
	splits := strings.Split(filter, ".")
	field, value := splits[0], splits[1]
	if field == "location" || field == "category" {
		return map[string]string{field: value}
	} else {
		return map[string]string{}
	}
}
func filterMapToDbReadableValues(filterMap map[string]string, keys []string) map[string]string {
	// creates a new map from previous filterMap to change key values acceptable for db query
	filterMapDb := map[string]string{}
	for k, v := range filterMap {
		if k == "location" {
			filterMapDb["l.name"] = v
		} else if k == "category" {
			filterMapDb["jc.title"] = v
		}

	}
	return filterMapDb
}

// ApplyFilters applies a set of filters to a record list.
// Each record will be checked against each filter.
// The filters are applied in the order they are passed in.
//func ApplyFilters(job *[]Job, filters ...Filter) []Job {
//	// Make sure there are actually filters to be applied.
//	if len(filters) == 0 {
//		return *job
//	}
//	filteredJobs := make([]Job, 0, len(*job))
//	// Range over the records and apply all the filters to each record.
//	// If the record passes all the filters, add it to the final slice.
//	for _, r := range *job {
//		keep := true
//
//		for _, f := range filters {
//			if !f(r) {
//				keep = false
//				break
//			}
//		}
//
//		if keep {
//			filteredJobs = append(filteredJobs, r)
//		}
//	}
//	return filteredJobs
//}

//func (j []Job) ApplyFilters() *Job {
//	if len(filters) == 0 {
//		return &j
//	}
//	filteredRecords := make([]string, 0, len(&j))
//	// Range over the records and apply all the filters to each record.
//	// If the record passes all the filters, add it to the final slice.
//	for _, r := range j {
//		keep := true
//
//		for _, f := range filters {
//			if !f(r) {
//				keep = false
//				break
//			}
//		}
//
//		if keep {
//			filteredRecords = append(filteredRecords, r)
//		}
//	}
//
//	return filteredRecords
//}
