package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
	"time"
)

type AccountService struct {
	Repository models.AccountRepository
	Logger     utils.Logger
	Timeout    time.Duration
	Config     models.Config
}

func NewAccountService(repository models.AccountRepository, logger utils.Logger, timeout time.Duration, cfg models.Config) *AccountService {
	return &AccountService{
		Repository: repository,
		Logger:     logger,
		Timeout:    timeout,
		Config:     cfg,
	}
}

const errorCreatingAccountErrMsg = "Error creating user"

func (s *AccountService) CreateAccount(c context.Context, req *models.Account) error {
	logger := s.Logger.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	account := &models.Account{
		Id:          req.Id,
		UserID:      req.UserID,
		AccountType: req.AccountType,
		Balance:     req.Balance,
	}

	_, err := s.Repository.CreateAccount(ctx, account)
	if err != nil {
		logger.Error("Error creating account", zap.Error(err))
		return errors.Wrap(err, errorCreatingAccountErrMsg)
	}

	return nil
}

func (s *AccountService) GetAccount(c context.Context, id uuid.UUID) (*models.Account, error) {
	logger := s.Logger.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	account, err := s.Repository.GetAccountById(ctx, id)
	if err != nil {
		logger.Error("Error getting account", zap.Error(err))
		return nil, err
	}

	return account, nil
}

func (s *AccountService) GetAllAccounts(c context.Context) ([]*models.Account, error) {
	logger := s.Logger.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	accounts, err := s.Repository.GetAllAccounts(ctx)
	if err != nil {
		logger.Error("Error getting all accounts", zap.Error(err))
		return nil, err
	}

	return accounts, nil
}
