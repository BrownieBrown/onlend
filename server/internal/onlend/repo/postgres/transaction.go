package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"server/pkg/models"
)

func NewTransactionRepository(db *sql.DB) models.TransactionRepository {
	return &Repository{db: db}
}

func (r *Repository) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	transaction.Id = uuid.New()

	query := "INSERT INTO transactions (id, sender_id, receiver_id, amount, transaction_type, status) VALUES ($1, $2, $3, $4, $5, $6) returning id"
	var returnedId uuid.UUID
	err := r.db.QueryRowContext(ctx, query, transaction.Id, transaction.SenderID, transaction.ReceiverID, transaction.Amount, transaction.TransactionType, transaction.Status).Scan(&returnedId)
	if err != nil {
		return nil, err
	}

	transaction.Id = returnedId
	return transaction, nil
}

func (r *Repository) GetTransaction(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	query := "SELECT id, sender_id, receiver_id, amount, transaction_type, status FROM transactions WHERE id = $1"
	var transaction models.Transaction
	err := r.db.QueryRowContext(ctx, query, id).Scan(&transaction.Id, &transaction.SenderID, &transaction.ReceiverID, &transaction.Amount, &transaction.TransactionType, &transaction.Status)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *Repository) GetAllTransaction(ctx context.Context) ([]*models.Transaction, error) {
	query := "SELECT id, sender_id, receiver_id, amount, transaction_type, status FROM transactions"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err = rows.Scan(&transaction.Id, &transaction.SenderID, &transaction.ReceiverID, &transaction.Amount, &transaction.TransactionType, &transaction.Status)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}
