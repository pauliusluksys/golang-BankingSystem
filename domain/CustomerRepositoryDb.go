package domain

import (
	"bankingV2/errs"
	"bankingV2/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type CustomerRepositoryDb struct {
	Client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(Id string) (*Customer, *errs.AppError) {
	//CustomerRep := NewCustomerRepositoryDb()
	var c Customer
	findByIdSql := "select customer_id, name,city,zipcode,date_of_birth,status from customers where customer_id=?"
	err := d.Client.Get(&c, findByIdSql, Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while querying customer table" + err.Error())
			return nil, errs.NewUnexpectError("Unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(UrlStatus string) ([]Customer, *errs.AppError) {

	customers := make([]Customer, 0)
	var statusNum int
	var queryByStatus bool
	if UrlStatus == "active" {
		statusNum = 1
		queryByStatus = true
		fmt.Println(statusNum)
		customers, _ = findByStatus(statusNum, queryByStatus, customers, d.Client)
	} else if UrlStatus == "inactive" {
		statusNum = 0
		queryByStatus = true
		fmt.Println(statusNum, queryByStatus, customers)

		customers, _ = findByStatus(statusNum, queryByStatus, customers, d.Client)
	} else {
		return findByStatus(statusNum, queryByStatus, customers, d.Client)
	}
	return customers, nil
}
func findByStatus(stat int, queryByStat bool, cust []Customer, client *sqlx.DB) ([]Customer, *errs.AppError) {
	cust = nil
	CustomerRep := NewCustomerRepositoryDb(client)
	var findAllSql string
	var err error
	if queryByStat == true {
		findAllSql = "select customer_id, name,city,zipcode,date_of_birth,status from customers where status=?"
		err = CustomerRep.Client.Select(&cust, findAllSql, stat)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("status not found")
			} else {
				log.Println("Error while querying customer table" + err.Error())
				return nil, errs.NewNotFoundError("Unexpected database error")
			}
		}
	} else {
		findAllSql = "select customer_id, name,city,zipcode,date_of_birth,status from customers"
		err = CustomerRep.Client.Select(&cust, findAllSql)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("status not found")
			} else {
				logger.Error("Error while querying customer table" + err.Error())
				return nil, errs.NewNotFoundError("Unexpected database error")
			}
		}
	}

	if err != nil {
		logger.Error("Error while scanning customers" + err.Error())
		return nil, errs.NewNotFoundError("Unexpected database error")
	}
	return cust, nil
}
func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{dbClient}
}
