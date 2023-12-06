package models

import "github.com/google/uuid"

type Transaction struct {
	Id              uuid.UUID       `json:"id" db:"id"`
	SenderID        uuid.UUID       `json:"senderId" db:"senderId"`
	ReceiverID      uuid.UUID       `json:"receiverId" db:"receiverId"`
	Amount          float64         `json:"amount" db:"amount"`
	TransactionType TransactionType `json:"transactionType" db:"transactionType"`
	Status          Status          `json:"status" db:"status"`
}

type TransactionService interface {
	CreateTransaction(transaction *Transaction) (uint, error)
	GetTransaction(id uint) (*Transaction, error)
	GetAllTransaction() ([]*Transaction, error)
	DeleteTransaction(id uint) error
}
