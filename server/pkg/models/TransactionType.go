package models

type TransactionType int

const (
	Send TransactionType = iota
	Receive
)
