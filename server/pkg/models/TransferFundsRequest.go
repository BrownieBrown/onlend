package models

import "github.com/google/uuid"

type TransferFundsRequest struct {
	SenderID   uuid.UUID `json:"senderId"`
	ReceiverID uuid.UUID `json:"receiverId"`
	Amount     float64   `json:"amount"`
}
