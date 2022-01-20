package domain

import (
	"bankingV2/dto"
	"bankingV2/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}
func (a Account) CanWithdraw(amount float64) bool {
	if amount > a.Amount {
		return false
	} else {
		return true
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	ById(Id string) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
}
