package domain

import (
	"bankingV2/dto"
	"bankingV2/errs"
	"gorm.io/gorm"
	"log"
)

type JobGorm struct {
	gorm.Model
	Title         string
	Description   string
	ApplyUntil    string
	LocationId    string
	JobCategoryId string
	IsPublished   string
	WhenToPublish string
	CreatedById   string
}
type City struct {
	ID      uint
	Name    string
	JobGorm JobGorm
}
type JobCategory struct {
	ID      uint
	Title   string
	JobGorm JobGorm
}

func (JobGorm) TableName() string {
	return "available_jobs"
}

type JobRepositoryDbGorm struct {
	Client *gorm.DB
}
type JobGormRepo interface {
	Save(job *JobGorm) (JobGorm, *errs.AppError)
}

func (d JobRepositoryDbGorm) Save(job *JobGorm) (*JobGorm, *errs.AppError) {
	result := d.Client.Model(&job).Updates(job)
	if result.Error != nil {
		log.Println("Error with database:", result.Error)
		return nil, errs.GormQueryError("Unexpected database error")
	}
	return job, nil
}

func (d JobRepositoryDbGorm) DeleteJob(job *JobGorm) (int64, *errs.AppError) {
	result := d.Client.Unscoped().Delete(&JobGorm{}, job.Model.ID)
	if result.Error != nil {
		log.Println("Error with database:", result.Error)
		return 0, errs.GormQueryError("Unexpected database error")
	}
	return result.RowsAffected, nil
}

func (d JobRepositoryDbGorm) UpdateJob(job *JobGorm) (*JobGorm, *errs.AppError) {
	result := d.Client.Create(job)
	if result.Error != nil {
		log.Println("Error with database:", result.Error)
		return nil, errs.GormQueryError("Unexpected database error")
	}
	return job, nil
}

func NewJobRepositoryDbGorm(dbClient *gorm.DB) JobRepositoryDbGorm {
	return JobRepositoryDbGorm{dbClient}
}

func (job JobGorm) ToJobResponseGormDto() dto.JobResponseGorm {

	return dto.JobResponseGorm{
		ID: job.ID,
	}
}
