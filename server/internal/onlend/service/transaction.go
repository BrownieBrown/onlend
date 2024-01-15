package service

import (
	"context"
	"github.com/google/uuid"
	"server/pkg/models"
)

type TransactionService struct {
	Repository     models.TransactionRepository
	AccountService models.AccountService
	Config         models.Config
}

func NewTransactionService(repository models.TransactionRepository, accountService models.AccountService, cfg models.Config) *TransactionService {
	return &TransactionService{
		Repository:     repository,
		AccountService: accountService,
		Config:         cfg,
	}
}

func (ts *TransactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	response, err := ts.Repository.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ts *TransactionService) GetTransaction(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	response, err := ts.Repository.GetTransaction(ctx, id)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ts *TransactionService) GetAllTransactions(ctx context.Context) ([]*models.Transaction, error) {
	response, err := ts.Repository.GetAllTransaction(ctx)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ts *TransactionService) TransferFunds(ctx context.Context, sender, receiver uuid.UUID, sum float64) ([]*models.Account, error) {
	tx, err := ts.Repository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	senderAccount, err := ts.updateSenderAccount(ctx, sender, sum)
	if err != nil {
		return nil, err
	}

	transActionSend := &models.Transaction{
		SenderID:        senderAccount.UserID,
		ReceiverID:      receiver,
		Amount:          sum,
		TransactionType: models.Send,
		Status:          models.Pending,
	}

	_, err = ts.Repository.CreateTransaction(ctx, transActionSend)

	receiverAccount, err := ts.updateReceiverAccount(ctx, receiver, sum)
	if err != nil {
		return nil, err
	}

	transActionReceive := &models.Transaction{
		SenderID:        sender,
		ReceiverID:      receiverAccount.UserID,
		Amount:          sum,
		TransactionType: models.Receive,
		Status:          models.Pending,
	}

	_, err = ts.Repository.CreateTransaction(ctx, transActionReceive)

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return []*models.Account{senderAccount, receiverAccount}, nil
}

func (ts *TransactionService) updateSenderAccount(c context.Context, senderId uuid.UUID, sum float64) (*models.Account, error) {
	account, err := ts.AccountService.GetAccountByUserId(c, senderId)
	if err != nil {
		return nil, err
	}
	return ts.AccountService.UpdateAccount(c, account.Id, sum, models.Send)
}
func (ts *TransactionService) updateReceiverAccount(c context.Context, receiverId uuid.UUID, sum float64) (*models.Account, error) {
	account, err := ts.AccountService.GetAccountByUserId(c, receiverId)
	if err != nil {
		return nil, err
	}
	return ts.AccountService.UpdateAccount(c, account.Id, sum, models.Receive)
}
