package model

type CreateAccountParams struct {
	AccountName string `json:"account_name" validate:"required"`
}

type CreateAccountResponse struct {
	AccountNo   string `json:"account_no"`
	AccountName string `json:"account_name"`
}

type GetAccountParams struct {
	AccountNo string `json:"account_no" validate:"required"`
}

type GetAccountResponse struct {
	AccountNo   string  `json:"account_no"`
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}
