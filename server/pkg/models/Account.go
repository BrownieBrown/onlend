package models

import (
	"context"
	"github.com/google/uuid"
)

type Account struct {
	Id          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"userId" db:"user_id"`
	AccountType string    `json:"accountType" db:"account_type"`
	Balance     float64   `json:"balance" db:"balance"`
}

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *Account) (*Account, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (*Account, error)
	GetAllAccounts(ctx context.Context) ([]*Account, error)
}

type AccountService interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccount(ctx context.Context, id uuid.UUID) (*Account, error)
	GetAllAccounts(ctx context.Context) ([]*Account, error)
}
