package domain

import (
	"bankingV2/errs"
	"bankingV2/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type AccountRepositoryDb struct {
	Client *sqlx.DB
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.Client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")
	}
	result, _ := tx.Exec("INSERT INTO transactions (account_id,amount,transaction_type,transaction_date)  values (?,?,?,?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec(`Update accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`Update accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}
	if err != nil {
		tx.Rollback()
		logger.Error("error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")
	}
	//getting the last transaction ID from transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected database error")
	}
	log.Println("it executes to this point!!!!")
	account, appErr := d.ById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}
func (d AccountRepositoryDb) ById(Id string) (*Account, *errs.AppError) {
	//CustomerRep := NewCustomerRepositoryDb()
	log.Println("it reaches this point")
	var a Account
	findByIdSql := "select account_id, customer_id,opening_date,account_type,amount,status from accounts where account_id=?"
	err := d.Client.Get(&a, findByIdSql, Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while querying customer table" + err.Error())
			return nil, errs.NewUnexpectError("Unexpected database error")
		}
	}
	return &a, nil
}
func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	log.Println("////////////////////////////////////////////////////")
	log.Println(a.CustomerId)
	sqlInsert := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) values (?,?,?,?,?)"

	result, err := d.Client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
