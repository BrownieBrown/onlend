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

func TestGetTransaction(t *testing.T) {
	mockTransactionService, handler, ctrl := setupTransactionHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()
	testUUID := uuid.New()
	senderId := uuid.New()
	receiverId := uuid.New()
	sum := 1000.0
	target := "/api/v1/transactions/" + testUUID.String()

	c, rec := prepareTestRequest(e, "", target, http.MethodGet)
	c.SetParamNames("id")
	c.SetParamValues(testUUID.String())

	transaction := helpers.CreateTransaction(testUUID, senderId, receiverId, sum, models.Send, models.Pending)
	mockTransactionService.EXPECT().GetTransaction(gomock.Any(), testUUID).Return(transaction, nil)

	err := handler.GetTransaction(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetAllTransactions(t *testing.T) {
	mockTransactionService, handler, ctrl := setupTransactionHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()
	c, rec := prepareTestRequest(e, "", "/api/v1/transactions", http.MethodGet)

	transaction := helpers.CreateTransaction(uuid.New(), uuid.New(), uuid.New(), 1000, models.Send, models.Pending)
	mockTransactionService.EXPECT().GetAllTransactions(gomock.Any()).Return([]*models.Transaction{transaction}, nil)

	err := handler.GetAllTransactions(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestTransferFunds(t *testing.T) {
	mockTransactionService, handler, ctrl := setupTransactionHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	senderId := uuid.New()
	receiverId := uuid.New()
	sum := 1000.0
	senderAccount := helpers.CreateAccount(uuid.New(), senderId, "default", 1000)
	receiverAccount := helpers.CreateAccount(uuid.New(), receiverId, "default", 1000)
	transferFundsRequest := helpers.CreateTransferFundsRequest(senderId, receiverId, 1000)

	e := echo.New()
	c, rec := prepareTestRequest(e, transferFundsRequest, "/api/v1/transactions", http.MethodPost)

	mockTransactionService.EXPECT().TransferFunds(gomock.Any(), senderId, receiverId, sum).Return([]*models.Account{senderAccount, receiverAccount}, nil)

	err := handler.TransferFunds(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func setupTransactionHandler(t *testing.T) (*mocks.MockTransactionService, *rest.TransactionHandler, *gomock.Controller) {
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
	mockTransactionService := mocks.NewMockTransactionService(ctrl)
	transactionHandler := rest.NewTransactionHandler(mockTransactionService, logger, cfg)

	return mockTransactionService, transactionHandler, ctrl
}
