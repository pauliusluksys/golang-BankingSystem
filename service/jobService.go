package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/errs"
	"log"
	"time"
)

type JobService interface {
	GetAllJobs(req dto.AllJobsRequest, filterMap map[string]string) (*[]dto.JobResponse, *errs.AppError)
	GetById(string) (*dto.JobResponse, *errs.AppError)
	NewJob(req dto.NewJobRequest) (*dto.JobResponse, *errs.AppError)
}
type DefaultJobService struct {
	repo domain.JobRepositoryDb
}

func (j DefaultJobService) NewJob(req dto.NewJobRequest) (*dto.JobResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err

	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	log.Println(a)
	NewAccount, err := s.repo.Save(a)
	if err != nil {
		log.Println("SOMETHING WENT WRONG")
		return nil, err
	}
	response := NewAccount.ToNewAccountResponseDto()
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
