package model

import "time"

type GetTransactionParams struct {
	AccountNo string `json:"account_no" validate:"required"`
}

type Transaction struct {
	ID              int64     `json:"id"`
	FromAccountNO   string    `json:"from_account_no"`
	TOAccountNO     string    `json:"to_account_no"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	TransactionTime time.Time `json:"transaction_time"`
}
