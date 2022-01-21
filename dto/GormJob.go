package dto

import "bankingV2/errs"

type JobRequestGorm struct {
	ID             uint   `json:"id"`
	CreatedById    string `json:"created_by_id"`
	Title          string `json:"job_title"`
	Description    string `json:"job_description"`
	CategoryNameId string `json:"job_category_id"`
	ApplyUntil     string `json:"apply_until"`
	CityId         string `json:"city_id"`
	IsPublished    string `json:"is_published"`
	WhenToPublish  string `json:"when_to_publish"`
}
type JobResponseGorm struct {
	ID uint `json:"job_id"`
}

func (newJob JobRequestGorm) Validate() *errs.AppError {
	return nil
}
