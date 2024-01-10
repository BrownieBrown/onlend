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
	"time"
)

func TestCreateAccount(t *testing.T) {
	accountService, mockRepository, ctrl, _ := setupMockAccountService(t)
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

func TestGetAccountByUUI(t *testing.T) {
	accountService, mockRepository, ctrl, _ := setupMockAccountService(t)
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

func TestGetAllAccounts(t *testing.T) {
	accountService, mockRepository, ctrl, _ := setupMockAccountService(t)
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

func setupMockAccountService(t *testing.T) (*service.AccountService, *mocks.MockAccountRepository, *gomock.Controller, utils.Logger) {
	utils.SetEnvVars()
	ctrl := gomock.NewController(t)

	logger, err := utils.NewZapLogger()
	if err != nil {
		t.Fatalf("Failed to load logger: %v", err)
	}

	mockRepository := mocks.NewMockAccountRepository(ctrl)
	accountService := service.NewAccountService(mockRepository, logger, time.Second, models.Config{})
	return accountService, mockRepository, ctrl, logger
}
