package models

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID       `json:"id" db:"id"`
	SenderID        uuid.UUID       `json:"senderId" db:"senderId"`
	ReceiverID      uuid.UUID       `json:"receiverId" db:"receiverId"`
	Amount          float64         `json:"amount" db:"amount"`
	TransactionType TransactionType `json:"transactionType" db:"transactionType"`
	Status          Status          `json:"status" db:"status"`
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error)
	GetTransaction(ctx context.Context, id uuid.UUID) (*Transaction, error)
	GetAllTransaction(ctx context.Context) ([]*Transaction, error)
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type TransactionService interface {
	TransferFunds(ctx context.Context, sender, receiver uuid.UUID, sum float64) ([]*Account, error)
	GetTransaction(ctx context.Context, id uuid.UUID) (*Transaction, error)
	GetAllTransactions(ctx context.Context) ([]*Transaction, error)
	CreateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error)
}
