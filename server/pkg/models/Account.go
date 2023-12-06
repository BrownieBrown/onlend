package models

import "github.com/google/uuid"

type Account struct {
	Id          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"userId" db:"userId"`
	AccountType string    `json:"accountType" db:"accountType"`
	Balance     float64   `json:"balance" db:"balance"`
}

type AccountService interface {
	CreateAccount(account *Account) (uint, error)
	GetAccount(id uint) (*Account, error)
	GetAllAccounts() ([]*Account, error)
	DeleteAccount(id uint) error
}
