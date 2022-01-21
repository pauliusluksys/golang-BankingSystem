package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/errs"
	"gorm.io/gorm"
	"log"
)

type JobService interface {
	GetAllJobs(req dto.AllJobsRequest, filterMap map[string]string) (*[]dto.JobResponse, *errs.AppError)
	GetById(string) (*dto.JobResponse, *errs.AppError)
}
type JobServiceGorm interface {
	NewJob(req dto.JobRequestGorm) (*dto.JobResponseGorm, *errs.AppError)
	UpdateJob(req dto.JobRequestGorm) (*dto.JobResponseGorm, *errs.AppError)
	DeleteJob(req dto.JobRequestGorm) (string, *errs.AppError)
}
type DefaultJobService struct {
	repo domain.JobRepositoryDb
}
type DefaultJobServiceGorm struct {
	repo domain.JobRepositoryDbGorm
}

func (j DefaultJobServiceGorm) DeleteJob(req dto.JobRequestGorm) (string, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return "", err

	}
	a := domain.JobGorm{
		Model: gorm.Model{ID: req.ID},
	}
	a.TableName()
	log.Println(a)
	result, err := j.repo.DeleteJob(&a)
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return "", err
	}
	var responseMSG string
	if result == 0 {
		responseMSG = "did not delete any rows, check if provided Id is right"
	} else {
		responseMSG = "deleted successfully"
	}
	return responseMSG, nil
}
func (j DefaultJobServiceGorm) UpdateJob(req dto.JobRequestGorm) (*dto.JobResponseGorm, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err

	}
	a := domain.JobGorm{
		Model:         gorm.Model{ID: req.ID},
		Title:         req.Title,
		Description:   req.Description,
		ApplyUntil:    req.ApplyUntil,
		JobCategoryId: req.CategoryNameId,
		LocationId:    req.CityId,
		IsPublished:   req.IsPublished,
		WhenToPublish: req.WhenToPublish,
	}
	a.TableName()
	log.Println(a)
	NewJobGorm, err := j.repo.Save(&a)
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err
	}
	response := NewJobGorm.ToJobResponseGormDto()
	return &response, nil
}
func (j DefaultJobServiceGorm) NewJob(req dto.JobRequestGorm) (*dto.JobResponseGorm, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err

	}

	a := domain.JobGorm{
		Title:         req.Title,
		Description:   req.Description,
		ApplyUntil:    req.ApplyUntil,
		JobCategoryId: req.CategoryNameId,
		LocationId:    req.CityId,
		IsPublished:   req.IsPublished,
		WhenToPublish: req.WhenToPublish,
		CreatedById:   req.CreatedById,
	}
	a.TableName()
	log.Println(a)
	NewJobGorm, err := j.repo.Save(&a)
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err
	}
	response := NewJobGorm.ToJobResponseGormDto()
	return &response, nil
}

func (d DefaultJobService) GetById(id string) (*dto.JobResponse, *errs.AppError) {
	j, err := d.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := j.ToDto()

	return &response, nil
}
func (d DefaultJobService) GetAllJobs(req dto.AllJobsRequest, filterMap map[string]string) (*[]dto.JobResponse, *errs.AppError) {

	j, err := d.repo.FindAll(filterMap)
	var jResponse []dto.JobResponse
	if err != nil {
		return nil, err
	}

	for _, value := range j {
		if req.Location != "" && req.CategoryName != "" {
			//value.ApplyFilters(req.Location, req.CategoryName)

		}
		jResponse = append(jResponse, value.ToDto())
	}
	return &jResponse, nil
}

//func (d DefaultJobService) GetById(a string) (*dto.JobResponse, *errs.AppError) {
//	c, err := d.repo.ById(id)
//	if err != nil {
//		return nil, err
//	}
//	response := d.ToDto()
//
//	return &response, nil
//}
func NewJobService(repo domain.JobRepositoryDb) DefaultJobService {
	return DefaultJobService{repo}
}
func NewJobServiceGorm(repo domain.JobRepositoryDbGorm) DefaultJobServiceGorm {
	return DefaultJobServiceGorm{repo}
}
