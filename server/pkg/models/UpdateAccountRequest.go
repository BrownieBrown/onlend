package models

type UpdateAccountRequest struct {
	Id              string          `json:"id"`
	Sum             float64         `json:"sum"`
	TransactionType TransactionType `json:"transactionType"`
}
