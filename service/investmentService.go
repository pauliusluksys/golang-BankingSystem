package service

import (
	"bankingV2/domain"
	"bankingV2/dto"
	"bankingV2/dto/DtoInvestment"
	"bankingV2/errs"
	"database/sql"
	"time"
)

type InvestmentService interface {
	CreateCustomerInvestment(request DtoInvestment.NewCustomerInvestmentRequest) (*DtoInvestment.NewCustomerInvestmentResponse, *errs.AppError)
	GetAllCustomerInvestments(customerID string) (*DtoInvestment.AllCustomerInvestmentResponse, *errs.AppError)
	//GetAllCustomersInvestments() (*DtoInvestment.AllCustomersResponse, *errs.AppError)
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

//func (S DefaultInvestmentService) GetAllCustomersInvestments() (*DtoInvestment.AllCustomersResponse, *errs.AppError) {
//	customersInvestments, err := S.repo.FindAllCustomerInvestments()
//	if err != nil {
//		return nil, err
//	} else {
//
//		for key, index := range customersInvestments {
//			customersInvestments.
//		}
//
//		return DtoInvestment.AllCustomersResponse, nil
//	}
//
//}
func (S DefaultInvestmentService) GetAllCustomerInvestments(customerID string) (*DtoInvestment.AllCustomerInvestmentResponse, *errs.AppError) {
	customerInvestments, err := S.repo.FindAllInvestmentsByCustomerId(customerID)
	customerInvestmentsCount, err := S.repo.FindAllCustomerInvestmentsCount(customerID)
	var cIResponseSlice []DtoInvestment.CustomerInvestmentResponse
	if err != nil {
		return nil, err
	} else {
		for _, cI := range customerInvestments {

			cIResponse := cI.CustomerInvestmentsToDto()
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
		AmountInvested:              request.AmountInvested,
		IsWithdrawn:                 request.IsWithdrawn,
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
