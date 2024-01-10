package models

type CreateAccountResponse struct {
	Id          string  `json:"id"`
	UserId      string  `json:"userId"`
	AccountType string  `json:"accountType"`
	Balance     float64 `json:"balance"`
}
