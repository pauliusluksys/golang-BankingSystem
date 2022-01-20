package dto

type NewTransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	AccountAmount float64 `json:"amount"`
}
