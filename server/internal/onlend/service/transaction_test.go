package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"server/internal/onlend/service"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"testing"
)

func TestGetTransaction(t *testing.T) {
	transactionService, _, mockTransactionRepo, ctrl := setupMockTransactionService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	transaction := &models.Transaction{
		Id:              uuid.New(),
		SenderID:        uuid.New(),
		ReceiverID:      uuid.New(),
		Amount:          100,
		TransactionType: models.Send,
	}

	mockTransactionRepo.EXPECT().GetTransaction(ctx, transaction.Id).Return(transaction, nil)
	resp, err := transactionService.GetTransaction(ctx, transaction.Id)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, transaction, resp)
}

func TestGetAllTransactions(t *testing.T) {
	transactionService, _, mockTransactionRepo, ctrl := setupMockTransactionService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	transaction := &models.Transaction{
		Id:              uuid.New(),
		SenderID:        uuid.New(),
		ReceiverID:      uuid.New(),
		Amount:          100,
		TransactionType: models.Send,
	}

	mockTransactionRepo.EXPECT().GetAllTransaction(ctx).Return([]*models.Transaction{transaction}, nil)
	resp, err := transactionService.GetAllTransactions(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, []*models.Transaction{transaction}, resp)
}

func TestCreateTransaction(t *testing.T) {
	transactionService, _, mockTransactionRepo, ctrl := setupMockTransactionService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	transaction := &models.Transaction{
		SenderID:        uuid.New(),
		ReceiverID:      uuid.New(),
		Amount:          100,
		TransactionType: models.Send,
	}

	mockTransactionRepo.EXPECT().CreateTransaction(ctx, transaction).Return(transaction, nil)
	resp, err := transactionService.CreateTransaction(ctx, transaction)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, transaction, resp)
}

func setupMockTransactionService(t *testing.T) (*service.TransactionService, *mocks.MockAccountRepository, *mocks.MockTransactionRepository, *gomock.Controller) {
	utils.UnsetEnvVars()
	ctrl := gomock.NewController(t)

	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)
	mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
	accountService := service.NewAccountService(mockAccountRepo, models.Config{})
	transactionService := service.NewTransactionService(mockTransactionRepo, accountService, models.Config{})
	return transactionService, mockAccountRepo, mockTransactionRepo, ctrl
}
