package postgres

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
)

func NewAccountRepository(db DBTX, logger utils.Logger) models.AccountRepository {
	return &repository{db: db, logger: logger}
}

func (r *repository) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	logger := r.logger.GetLogger()

	query := "INSERT INTO accounts (id, user_id, account_type, balance) VALUES ($1, $2, $3, $4) returning id"
	var returnedId uuid.UUID
	err := r.db.QueryRowContext(ctx, query, account.Id, account.UserID, account.AccountType, account.Balance).Scan(&returnedId)
	if err != nil {
		logger.Error("Error while creating account", zap.Error(err))
		return nil, err
	}

	account.Id = returnedId
	return account, nil
}

func (r *repository) GetAccountById(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	logger := r.logger.GetLogger()

	query := "SELECT id, user_id, account_type, balance FROM accounts WHERE id = $1"
	var account models.Account
	err := r.db.QueryRowContext(ctx, query, id).Scan(&account.Id, &account.UserID, &account.AccountType, &account.Balance)
	if err != nil {
		logger.Error("Error while finding account by id", zap.Error(err))
		return nil, err
	}

	return &account, nil
}

func (r *repository) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	logger := r.logger.GetLogger()

	query := "SELECT id, user_id, account_type, balance FROM accounts"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		logger.Error("Error while getting all accounts", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.Id, &account.UserID, &account.AccountType, &account.Balance)
		if err != nil {
			logger.Error("Error while scanning account", zap.Error(err))
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, nil
}
