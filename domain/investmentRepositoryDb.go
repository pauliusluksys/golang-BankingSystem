package domain

import (
	"bankingV2/dto"
	"bankingV2/dto/DtoInvestment"
	"bankingV2/errs"
	"bankingV2/logger"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type InvestmentRepositoryDbGorm struct {
	Client *gorm.DB
}
type InvestmentRepositoryDb struct {
	Client *sqlx.DB
}

const GormErrorUnique = "Record already exists, thus was not saved"

func (d InvestmentRepositoryDbGorm) FindAll() ([]InvestmentGorm, *errs.AppError) {
	var investment []InvestmentGorm
	result := d.Client.Find(&investment)
	if result.Error != nil {
		log.Println("Error with database:", result.Error)
		return nil, errs.GormQueryError("Unexpected database error")
	}
	return investment, nil
}
func (d InvestmentRepositoryDb) ById(Id uint) (*Investment, *errs.AppError) {

	var i Investment

	findByIdSql := "select * from investments where id=?"
	err := d.Client.Get(&i, findByIdSql, Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("investment not found")
		} else {
			log.Println("Error while querying investment table " + err.Error())
			return nil, errs.NewUnexpectError("Unexpected database error")
		}
	}
	return &i, nil
}
func (d InvestmentRepositoryDb) FindAllCustomerInvestmentsCount(customerID string) (*int, *errs.AppError) {
	var cICount []CustomerInvestmentsCount
	query := "SELECT COUNT(*) as total from investments ig inner join customer_investments ci on ig.id = ci.investment_id where ci.customer_id = ?;"
	err := d.Client.Select(&cICount, query, customerID)
	if err != nil {
		logger.Error("Error while selecting customer investments count" + err.Error())
		return nil, errs.NewUnexpectError("unexpected database error")
	} else {
		return &cICount[0].Total, nil
	}
}
func (d InvestmentRepositoryDb) FindAllCustomersInvestments(offset string, quantity string) ([]CustomerInvestment, *errs.AppError) {
	var customerInvestments []CustomerInvestment
	query := "SELECT ci.customer_id , c.name as customer_name , ci.invested_amount, ci.withdrawn_state,ci.created_at as customer_investment_created_at, ig.id as investment_id,ig.created_at as investment_created_at,ig.updated_at as investment_updated_at,ig.deleted_at as investment_deleted_at,ig.title as investment_title,ig.category_investment_id,ig.company_investment_id,ig.risk_level_investment_id,ci3.name as company_name, ci4.name as category_name, rl.name as risk_level_name from investments ig inner join customer_investment ci on ig.id = ci.investment_id inner join customers c on ci.customer_id = c.customer_id inner join company_investments ci3 ON ig.company_investment_id = ci3.id inner join category_investments ci4 ON ig.category_investment_id = ci4.id inner join risk_level_investments rl ON ig.risk_level_investment_id = rl.id ORDER BY customer_name,customer_id,investment_id asc limit ?,?;"
	err := d.Client.Select(&customerInvestments, query, offset, quantity)
	if err != nil {
		logger.Error("Error while selecting all customer investments: " + err.Error())
		return nil, errs.NewUnexpectError("unexpected database error")
	} else {
		fmt.Println("hello")
		fmt.Println(customerInvestments[0].InvestmentID, len(customerInvestments))
		return customerInvestments, nil
	}
}
func (d InvestmentRepositoryDb) FindAllInvestmentsByCustomerId(customerID string) ([]CustomerInvestment, *errs.AppError) {
	var customerInvestments []CustomerInvestment

	query := "SELECT ci.amount_invested, ci.is_withdrawn,ci.created_at as customer_investment_created_at, ig.id as investment_id,ig.created_at as investment_created_at,ig.updated_at as investment_updated_at,ig.deleted_at as investment_deleted_at,ig.title as investment_title,ig.category_investment_id,ig.company_investment_id,ig.risk_level_investment_id,ci3.name as company_name, ci4.name as category_name, rl.name as risk_level_name from investments ig inner join customer_investments ci on ig.id = ci.investment_id inner join company_investments ci3 ON ig.company_investment_id = ci3.id inner join category_investments ci4 ON ig.category_investment_id = ci4.id inner join risk_level_investments rl ON ig.risk_level_investment_id = rl.id where ci.customer_id = ?;"
	err := d.Client.Select(&customerInvestments, query, customerID)
	if err != nil {
		logger.Error("Error while selecting all customer investments by id: " + err.Error())
		return nil, errs.NewUnexpectError("unexpected database error")
	} else {
		return customerInvestments, nil
	}
}
func (d InvestmentRepositoryDb) CreateCustomerInvestment(ci CustomerInvestment) (*CustomerInvestment, *errs.AppError) {
	tx, err := d.Client.Begin()
	if err != nil {
		logger.Error("Error while creating customer investment: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")
	}
	_, err = tx.Exec("INSERT INTO customer_investments (customer_id,investment_id,amount_invested,is_withdrawn,created_at)  values (?,?,?,?,?)", ci.CustomerID, ci.InvestmentID, ci.InvestedAmount, ci.WithdrawnState, ci.CustomerInvestmentCreatedAt)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("failed to rollback after execution error: ", err)
		}
		logger.Error("error while saving customer investment: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")

	}
	err = tx.Commit()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			fmt.Println("failed to rollback after commit error: ", err)
			return nil, errs.NewUnexpectError("Unexpected database error")
		}
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")
	}
	dc := CustomerRepositoryDb{d.Client}
	di := InvestmentRepositoryDb{d.Client}
	customer, err1 := dc.ById(strconv.FormatUint(uint64(ci.CustomerID), 10))
	if err1 != nil {
		return nil, err1
	}
	investment, err2 := di.ById(ci.InvestmentID)
	if err2 != nil {
		return nil, err2
	}
	ci.CustomerName.String = customer.Name
	ci.InvestmentTitle = investment.Title
	return &ci, nil
}
func (d InvestmentRepositoryDbGorm) CreateRiskLevel(request DtoInvestment.NewRiskLevelRequest) (*RiskLevelInvestment, *errs.AppError) {
	var newRiskLevel RiskLevelInvestment
	newRiskLevel.Name = request.Name
	result := d.Client.Create(&newRiskLevel)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errs.GormQueryError(fmt.Sprintf("Whilst creating new investment got error: ", result.Error))
	} else if result.RowsAffected == 0 {
		return nil, errs.GormQueryError(fmt.Sprintf(GormErrorUnique, result.Error))
	}

	return &newRiskLevel, nil
}
func (d InvestmentRepositoryDbGorm) CreateCategory(request DtoInvestment.NewCategoryRequest) (*CategoryInvestment, *errs.AppError) {
	var newCategory CategoryInvestment
	newCategory.Name = request.Name
	result := d.Client.Create(&newCategory)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errs.GormQueryError(fmt.Sprintf("Whilst creating new investment got error: ", result.Error))
	} else if result.RowsAffected == 0 {
		return nil, errs.GormQueryError(fmt.Sprintf(GormErrorUnique))
	}
	return &newCategory, nil
}
func (d InvestmentRepositoryDbGorm) CreateCompany(request DtoInvestment.NewCompanyRequest) (*CompanyInvestment, *errs.AppError) {
	var newCompany CompanyInvestment
	newCompany.Name = request.Name
	result := d.Client.Create(&newCompany)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errs.GormQueryError(fmt.Sprintf("Whilst creating new investment got error: "))
	} else if result.RowsAffected == 0 {
		return nil, errs.GormQueryError(fmt.Sprintf(GormErrorUnique))
	}

	return &newCompany, nil
}
func (d InvestmentRepositoryDbGorm) CreateInvestment(request DtoInvestment.NewInvestmentRequest) (*InvestmentGorm, *errs.AppError) {
	var newInvestment InvestmentGorm
	newInvestment = requestToInvestment(request)
	result := d.Client.Create(&newInvestment)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errs.GormQueryError(fmt.Sprintf("Whilst creating new investment got error: ", result.Error))
	}
	return &newInvestment, nil
}
func NewInvestmentRepositoryDbGorm(dbClient *gorm.DB) InvestmentRepositoryDbGorm {
	return InvestmentRepositoryDbGorm{dbClient}
}
func NewInvestmentRepositoryDb(dbClient *sqlx.DB) InvestmentRepositoryDb {
	return InvestmentRepositoryDb{dbClient}
}

func (Investment InvestmentGorm) ToInvestmentResponseGormDto() dto.InvestmentResponseGorm {

	return dto.InvestmentResponseGorm{
		ID:    Investment.ID,
		Title: Investment.Title,
	}
}
func requestToInvestment(request DtoInvestment.NewInvestmentRequest) InvestmentGorm {
	return InvestmentGorm{
		Title:                 request.Title,
		CategoryInvestmentID:  request.InvestmentCategoryID,
		CompanyInvestmentID:   request.InvestmentCompanyID,
		RiskLevelInvestmentID: request.RiskLevelID,
	}
}
