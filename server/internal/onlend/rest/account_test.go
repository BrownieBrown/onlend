package rest_test

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"server/internal/helpers"
	"server/internal/onlend/rest"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"testing"
)

func TestGetAccount(t *testing.T) {
	mockAccountService, handler, ctrl := setupAccountHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()
	testUUID := uuid.New()
	testUserID := uuid.New()
	target := "/api/v1/accounts/" + testUUID.String()

	c, rec := prepareTestRequest(e, "", target, http.MethodGet)

	account := helpers.CreateAccount(testUUID, testUserID, "checking", 1000) // testUserID should be defined or relevant
	mockAccountService.EXPECT().GetAccount(gomock.Any(), testUUID).Return(account, nil)

	err := handler.GetAccount(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAllAccounts(t *testing.T) {
	mockAccountService, handler, ctrl := setupAccountHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()
	c, rec := prepareTestRequest(e, "", "/api/v1/accounts", http.MethodGet)

	account := helpers.CreateAccount(uuid.New(), uuid.New(), "checking", 1000)
	mockAccountService.EXPECT().GetAllAccounts(gomock.Any()).Return([]*models.Account{account}, nil)

	err := handler.GetAllAccounts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func setupAccountHandler(t *testing.T) (*mocks.MockAccountService, *rest.AccountHandler, *gomock.Controller) {
	utils.SetEnvVars()
	cfg, err := utils.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	ctrl := gomock.NewController(t)
	logger, err := utils.NewZapLogger()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	mockAccountService := mocks.NewMockAccountService(ctrl)
	handler := rest.NewAccountHandler(mockAccountService, logger, cfg)

	return mockAccountService, handler, ctrl
}
