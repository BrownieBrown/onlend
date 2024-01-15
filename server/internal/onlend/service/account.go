package service

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"server/pkg/models"
)

type AccountService struct {
	Repository models.AccountRepository
	Config     models.Config
}

func NewAccountService(repository models.AccountRepository, cfg models.Config) *AccountService {
	return &AccountService{
		Repository: repository,
		Config:     cfg,
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *models.Account) error {
	account := &models.Account{
		Id:          req.Id,
		UserID:      req.UserID,
		AccountType: req.AccountType,
		Balance:     req.Balance,
	}

	_, err := s.Repository.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	account, err := s.Repository.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) GetAccountByUserId(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	account, err := s.Repository.GetAccountByUserId(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}
	return account, nil
}

func (s *AccountService) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	accounts, err := s.Repository.GetAllAccounts(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	return accounts, nil
}

func (s *AccountService) UpdateAccount(ctx context.Context, id uuid.UUID, sum float64, transactionType models.TransactionType) (*models.Account, error) {
	account, err := s.Repository.GetAccountById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	newBalance := account.Balance
	switch transactionType {
	case models.Send:
		newBalance -= sum
		if newBalance < 0 {
			return nil, errors.New("Not enough money on account")
		}
	case models.Receive:
		newBalance += sum
	}

	updatedAccount, err := s.Repository.UpdateAccount(ctx, account.Id, newBalance)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}
