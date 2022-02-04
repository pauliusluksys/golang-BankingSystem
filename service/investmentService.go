package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/dto/DtoInvestment"
	"bankingV2/errs"
	"database/sql"
	"fmt"
	"time"
)

type InvestmentService interface {
	CreateCustomerInvestment(request DtoInvestment.NewCustomerInvestmentRequest) (*DtoInvestment.NewCustomerInvestmentResponse, *errs.AppError)
	GetAllCustomerInvestments(customerID string) (*DtoInvestment.AllCustomerInvestmentResponse, *errs.AppError)
	GetAllCustomersInvestments(offset string, quantity string) (*[]DtoInvestment.ByCustomerResponse, *errs.AppError)
}
type DefaultInvestmentService struct {
	repo domain.InvestmentRepositoryDb
}

func NewInvestmentService(repo domain.InvestmentRepositoryDb) DefaultInvestmentService {
	return DefaultInvestmentService{repo}
}

type InvestmentServiceGorm interface {
	GetAllInvestments() (*[]dto.InvestmentResponseGorm, *errs.AppError)
	CreateInvestment(request DtoInvestment.NewInvestmentRequest) (*DtoInvestment.NewInvestmentResponse, *errs.AppError)
	CreateRiskLevel(request DtoInvestment.NewRiskLevelRequest) (*DtoInvestment.NewRiskLevelResponse, *errs.AppError)
	CreateCompany(request DtoInvestment.NewCompanyRequest) (*DtoInvestment.NewCompanyResponse, *errs.AppError)
	CreateCategory(request DtoInvestment.NewCategoryRequest) (*DtoInvestment.NewCategoryResponse, *errs.AppError)
}

type DefaultInvestmentServiceGorm struct {
	repo domain.InvestmentRepositoryDbGorm
}

func NewInvestmentServiceGorm(repo domain.InvestmentRepositoryDbGorm) DefaultInvestmentServiceGorm {
	return DefaultInvestmentServiceGorm{repo}
}

func (S DefaultInvestmentService) GetAllCustomersInvestments(offset string, quantity string) (*[]DtoInvestment.ByCustomerResponse, *errs.AppError) {
	var cIByCustomer []DtoInvestment.ByCustomerResponse
	var cIByInvestment []DtoInvestment.ByInvestmentResponse
	var byInvestment DtoInvestment.ByInvestmentResponse
	var byCustomer DtoInvestment.ByCustomerResponse
	var cITotal = 0
	var fullAmountInvested = 0
	var cILastCreated time.Time

	customersInvestments, err := S.repo.FindAllCustomersInvestments(offset, quantity)

	var lastInvestmentKey = len(customersInvestments) - 1

	if err != nil {
		return nil, err
	} else {
		for key, index := range customersInvestments {
			cIResponse := index.CustomersInvestmentsToDto()
			fmt.Println(key)
			if key == 0 {
				cITotal = 1
				cILastCreated = index.CustomerInvestmentCreatedAt.Time
				fullAmountInvested = int(index.InvestedAmount)
				byIResponse := index.ByInvestmentToDto()

				byInvestment = byIResponse
				byInvestment.CustomerInvestmentResponse = append(byInvestment.CustomerInvestmentResponse, cIResponse)

			} else {
				if (customersInvestments[key].InvestmentID == customersInvestments[key-1].InvestmentID) && (customersInvestments[key].CustomerID == customersInvestments[key-1].CustomerID) {

					if cILastCreated.Before(index.CustomerInvestmentCreatedAt.Time) {
						cILastCreated = index.CustomerInvestmentCreatedAt.Time
					}
					cITotal++
					fullAmountInvested = fullAmountInvested + int(index.InvestedAmount)

					byInvestment.CustomerInvestmentResponse = append(byInvestment.CustomerInvestmentResponse, cIResponse)
				} else if (customersInvestments[key].InvestmentID != customersInvestments[key-1].InvestmentID) && (customersInvestments[key].CustomerID == customersInvestments[key-1].CustomerID) {
					fmt.Println("after else:", index.InvestmentID)
					cILastCreated = index.CustomerInvestmentCreatedAt.Time
					byInvestment.TotalInvestments = cITotal
					byInvestment.FullAmountInvested = fullAmountInvested
					cITotal = 1
					fullAmountInvested = int(index.InvestedAmount)

					cIByInvestment = append(cIByInvestment, byInvestment)
					byInvestment.CustomerInvestmentResponse = nil
					byIResponse := customersInvestments[key-1].ByInvestmentToDto()
					byInvestment = byIResponse
					byInvestment.CustomerInvestmentResponse = append(byInvestment.CustomerInvestmentResponse, cIResponse)
				} else if customersInvestments[key].CustomerID != customersInvestments[key-1].CustomerID {

					cILastCreated = index.CustomerInvestmentCreatedAt.Time
					byInvestment.TotalInvestments = cITotal
					byInvestment.FullAmountInvested = fullAmountInvested
					cITotal = 1
					fullAmountInvested = int(index.InvestedAmount)

					cIByInvestment = append(cIByInvestment, byInvestment)
					byCustomer = customersInvestments[key-1].ByCustomerToDto()
					byCustomer.ByInvestmentResponse = cIByInvestment
					cIByCustomer = append(cIByCustomer, byCustomer)
					cIByInvestment = nil
					byCustomer.ByInvestmentResponse = nil
					byInvestment.CustomerInvestmentResponse = nil

					byInvestment.CustomerInvestmentResponse = append(byInvestment.CustomerInvestmentResponse, cIResponse)
				}
			}
			byInvestment.CustomerInvestmentLastCreated = cILastCreated.String()
			if key == lastInvestmentKey {
				if !(customersInvestments[key].InvestmentID == customersInvestments[key-1].InvestmentID) && (customersInvestments[key].CustomerID == customersInvestments[key-1].CustomerID) {
					if cILastCreated.Before(index.CustomerInvestmentCreatedAt.Time) {
						cILastCreated = index.CustomerInvestmentCreatedAt.Time
						fullAmountInvested = fullAmountInvested + int(index.InvestedAmount)
					}
				}

				cILastCreated = index.CustomerInvestmentCreatedAt.Time
				byInvestment.TotalInvestments = cITotal
				cITotal = 1
				byInvestment.FullAmountInvested = fullAmountInvested
				cIByInvestment = append(cIByInvestment, byInvestment)
				byCustomer = index.ByCustomerToDto()
				byCustomer.ByInvestmentResponse = cIByInvestment
				cIByCustomer = append(cIByCustomer, byCustomer)

			}
		}
	}
	return &cIByCustomer, nil
}

func (S DefaultInvestmentService) GetAllCustomerInvestments(customerID string) (*DtoInvestment.AllCustomerInvestmentResponse, *errs.AppError) {
	customerInvestments, err := S.repo.FindAllInvestmentsByCustomerId(customerID)
	customerInvestmentsCount, err := S.repo.FindAllCustomerInvestmentsCount(customerID)
	var cIResponseSlice []DtoInvestment.CustomerInvestmentResponse
	if err != nil {
		return nil, err
	} else {
		for _, cI := range customerInvestments {

			cIResponse := cI.CustomersInvestmentsToDto()
			cIResponseSlice = append(cIResponseSlice, cIResponse)
		}
		AllCIResponse := DtoInvestment.AllCustomerInvestmentResponse{*customerInvestmentsCount, cIResponseSlice}
		return &AllCIResponse, nil
	}
}
func (s DefaultInvestmentService) CreateCustomerInvestment(request DtoInvestment.NewCustomerInvestmentRequest) (*DtoInvestment.NewCustomerInvestmentResponse, *errs.AppError) {
	timeVal := time.Now()
	customerInvestment := domain.CustomerInvestment{
		CustomerID:                  request.CustomerID,
		InvestmentID:                request.InvestmentID,
		InvestedAmount:              request.AmountInvested,
		WithdrawnState:              request.IsWithdrawn,
		CustomerInvestmentCreatedAt: sql.NullTime{timeVal, true},
	}
	ci, err := s.repo.CreateCustomerInvestment(customerInvestment)
	if err != nil {
		return nil, err
	} else {
		response := ci.NewCustomerInvestmentToDto()
		return &response, nil
	}

}
func (s DefaultInvestmentServiceGorm) CreateCompany(request DtoInvestment.NewCompanyRequest) (*DtoInvestment.NewCompanyResponse, *errs.AppError) {

	newCompany, err := s.repo.CreateCompany(request)
	if err != nil {
		return nil, err
	} else {
		response := newCompany.NewCompanyToDto()
		return &response, nil
	}
}
func (s DefaultInvestmentServiceGorm) CreateCategory(request DtoInvestment.NewCategoryRequest) (*DtoInvestment.NewCategoryResponse, *errs.AppError) {

	newCategory, err := s.repo.CreateCategory(request)
	if err != nil {
		return nil, err
	} else {
		response := newCategory.NewCategoryToDto()
		return &response, nil
	}
}
func (s DefaultInvestmentServiceGorm) CreateRiskLevel(request DtoInvestment.NewRiskLevelRequest) (*DtoInvestment.NewRiskLevelResponse, *errs.AppError) {

	newRiskLevel, err := s.repo.CreateRiskLevel(request)
	if err != nil {
		return nil, err
	} else {
		response := newRiskLevel.NewRiskLevelToDto()
		return &response, nil
	}
}
func (s DefaultInvestmentServiceGorm) CreateInvestment(request DtoInvestment.NewInvestmentRequest) (*DtoInvestment.NewInvestmentResponse, *errs.AppError) {
	err := request.Validate()
	if !(err == nil) {
		return nil, err
	} else {
		newInvestment, err := s.repo.CreateInvestment(request)
		if err != nil {
			return nil, err
		} else {
			response := newInvestment.NewInvestmentToDto()
			return &response, nil
		}
	}

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
