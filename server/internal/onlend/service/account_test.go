package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"server/internal/helpers"
	"server/internal/onlend/service"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	accountService, mockRepository, ctrl := setupMockAccountService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	req := &models.Account{
		Id:          uuid.New(),
		UserID:      uuid.New(),
		AccountType: "default",
		Balance:     0,
	}

	mockRepository.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(req, nil)
	err := accountService.CreateAccount(ctx, req)

	assert.NoError(t, err)
}

func TestGetAccountByUUID(t *testing.T) {
	accountService, mockRepository, ctrl := setupMockAccountService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	account := helpers.CreateAccount(uuid.New(), uuid.New(), "default", 0)

	mockRepository.EXPECT().GetAccountById(gomock.Any(), account.Id).Return(account, nil)
	resp, err := accountService.GetAccount(ctx, account.Id)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, account, resp)
}

func TestGetAccountByUserId(t *testing.T) {
	accountService, mockRepository, ctrl := setupMockAccountService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	account := helpers.CreateAccount(uuid.New(), uuid.New(), "default", 0)

	mockRepository.EXPECT().GetAccountByUserId(gomock.Any(), account.UserID).Return(account, nil)
	resp, err := accountService.GetAccountByUserId(ctx, account.UserID)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, account, resp)
}

func TestGetAllAccounts(t *testing.T) {
	accountService, mockRepository, ctrl := setupMockAccountService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	account := helpers.CreateAccount(uuid.New(), uuid.New(), "default", 0)
	accounts := []*models.Account{account}

	mockRepository.EXPECT().GetAllAccounts(gomock.Any()).Return(accounts, nil)
	resp, err := accountService.GetAllAccounts(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, accounts, resp)
}

func TestUpdateAccount(t *testing.T) {
	accountService, mockRepository, ctrl := setupMockAccountService(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	account := helpers.CreateAccount(uuid.New(), uuid.New(), "default", 0)

	mockRepository.EXPECT().GetAccountById(gomock.Any(), account.Id).Return(account, nil)
	mockRepository.EXPECT().UpdateAccount(gomock.Any(), account.Id, account.Balance).Return(account, nil)
	resp, err := accountService.UpdateAccount(ctx, account.Id, account.Balance, 1)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, account, resp)
}

func setupMockAccountService(t *testing.T) (*service.AccountService, *mocks.MockAccountRepository, *gomock.Controller) {
	utils.SetEnvVars()
	ctrl := gomock.NewController(t)

	mockRepository := mocks.NewMockAccountRepository(ctrl)
	accountService := service.NewAccountService(mockRepository, models.Config{})
	return accountService, mockRepository, ctrl
}
