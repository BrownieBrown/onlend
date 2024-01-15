package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"server/pkg/models"
)

func NewAccountRepository(db *sql.DB) models.AccountRepository {
	return &Repository{db: db}
}

func (r *Repository) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	query := "INSERT INTO accounts (id, user_id, account_type, balance) VALUES ($1, $2, $3, $4) returning id"
	var returnedId uuid.UUID
	err := r.db.QueryRowContext(ctx, query, account.Id, account.UserID, account.AccountType, account.Balance).Scan(&returnedId)
	if err != nil {
		return nil, err
	}

	account.Id = returnedId
	return account, nil
}

func (r *Repository) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	query := "SELECT id, user_id, account_type, balance FROM accounts WHERE id = $1"
	var account models.Account
	err := r.db.QueryRowContext(ctx, query, id).Scan(&account.Id, &account.UserID, &account.AccountType, &account.Balance)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *Repository) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	query := "SELECT id, user_id, account_type, balance FROM accounts"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.Id, &account.UserID, &account.AccountType, &account.Balance)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}

func (r *Repository) GetAccountByUserId(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	query := "SELECT id, user_id, account_type, balance FROM accounts WHERE user_id = $1"
	var account models.Account

	err := r.db.QueryRowContext(ctx, query, id).Scan(&account.Id, &account.UserID, &account.AccountType, &account.Balance)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *Repository) UpdateAccount(ctx context.Context, id uuid.UUID, balance float64) (*models.Account, error) {
	query := "UPDATE accounts SET balance = $1 WHERE id = $2 RETURNING id, user_id, account_type, balance"
	var account models.Account
	err := r.db.QueryRowContext(ctx, query, balance, id).Scan(&account.Id, &account.UserID, &account.AccountType, &account.Balance)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
