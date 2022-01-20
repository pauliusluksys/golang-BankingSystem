package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/errs"
)

type CustomerService interface {
	GetAllCustomer(string) (*[]dto.CustomerResponse, *errs.AppError)
	GetById(string) (*dto.CustomerResponse, *errs.AppError)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(UrlStatus string) (*[]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll(UrlStatus)
	var cResponse []dto.CustomerResponse
	if err != nil {
		return nil, err
	}
	for _, value := range c {
		cResponse = append(cResponse, value.ToDto())
	}
	return &cResponse, nil
}
func (s DefaultCustomerService) GetById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	return &response, nil
}
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
