package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/errs"
)

type InvestmentServiceGorm interface {
	GetAllInvestments() (*[]dto.InvestmentResponseGorm, *errs.AppError)
}
type DefaultInvestmentServiceGorm struct {
	repo domain.InvestmentRepositoryDbGorm
}

func NewInvestmentServiceGorm(repo domain.InvestmentRepositoryDbGorm) DefaultInvestmentServiceGorm {
	return DefaultInvestmentServiceGorm{repo}
}
func (s DefaultInvestmentServiceGorm) GetAllInvestments() (*[]dto.InvestmentResponseGorm, *errs.AppError) {
	i, err := s.repo.FindAll()
	var iResponse []dto.InvestmentResponseGorm
	if err != nil {
		return nil, err
	}
	for _, value := range i {
		iResponse = append(iResponse, value.ToInvestmentResponseGormDto())
	}
	return &iResponse, nil
}
