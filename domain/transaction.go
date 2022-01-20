package domain

import (
	"bankingV2/dto"
	"bankingV2/errs"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}
type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{TransactionId: t.TransactionId, AccountId: t.AccountId, AccountAmount: t.Amount, TransactionType: t.TransactionType, TransactionDate: t.TransactionDate}
}
func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == "withdrawal" {
		return true
	} else {
		return false
	}
}
