package models

type CreateTransactionRequest struct {
	SenderID        string          `json:"senderId"`
	ReceiverID      string          `json:"receiverId"`
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"transactionType"`
}
