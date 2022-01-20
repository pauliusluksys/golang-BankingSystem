package dto

type NewJobRequest struct {
	CreatedBy      string `json:"created_by_employee"`
	Title          string `json:"job_title"`
	Description    string `json:"job_description"`
	CategoryNameId string `json:"job_category"`
	ApplyUntil     string `json:"apply_until"`
	CityId         string `json:"city"`
	IsPublished    string `json:"is_published"`
	WhenToPublish  string `json:"when_to_publish"`
}

func (newJob NewJobRequest) validate() {
	
}
