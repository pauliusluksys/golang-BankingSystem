package dto

import (
	"bankingV2/errs"
	"strings"
)

type TransactionRequest struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string
}
type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	AccountAmount   float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0.00 {
		return errs.NewValidationError("Transaction amount cannot be negative")
	}
	if strings.ToLower(r.TransactionType) != "deposit" && !r.IsTransactionTypeWithdrawal() {
		return errs.NewValidationError("Transaction type can only be either deposit or withdrawal")
	}
	return nil
}
func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {

	return r.TransactionType == "withdrawal"
}
