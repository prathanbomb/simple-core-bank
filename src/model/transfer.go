package model

type TransferInParams struct {
	ToAccountNo string  `json:"to_account_no" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
}

type TransferInResponse struct {
	TransactionID int64   `json:"transaction_id"`
	ToAccountNo   string  `json:"to_account_no"`
	Amount        float64 `json:"amount"`
}
