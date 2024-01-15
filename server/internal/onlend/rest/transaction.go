package rest

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"server/internal/utils"
	"server/pkg/models"
)

type TransactionHandler struct {
	TransactionService models.TransactionService
	Logger             utils.Logger
	Config             models.Config
}

func NewTransactionHandler(ts models.TransactionService, logger utils.Logger, cfg models.Config) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: ts,
		Logger:             logger,
		Config:             cfg,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx echo.Context) error {
	logger := h.Logger.GetLogger()

	var req models.CreateTransactionRequest
	if err := ctx.Bind(&req); err != nil {
		logger.Error("failed to parse request", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.SenderID == "" || req.ReceiverID == "" || req.Amount <= 0 || req.TransactionType != models.TransactionType(models.Send) && req.TransactionType != models.TransactionType(models.Receive) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid or missing parameters"})
	}

	senderID, err := uuid.Parse(req.SenderID)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	receiverID, err := uuid.Parse(req.ReceiverID)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	transaction := &models.Transaction{
		SenderID:        senderID,
		ReceiverID:      receiverID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
	}

	response, err := h.TransactionService.CreateTransaction(ctx.Request().Context(), transaction)
	if err != nil {
		logger.Error("failed to create transaction", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create transaction"})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransaction(ctx echo.Context) error {
	logger := h.Logger.GetLogger()
	idStr := ctx.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	transaction, err := h.TransactionService.GetTransaction(ctx.Request().Context(), id)
	if err != nil {
		logger.Error("failed to get transaction", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get transaction"})
	}

	return ctx.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetAllTransactions(ctx echo.Context) error {
	logger := h.Logger.GetLogger()

	transactions, err := h.TransactionService.GetAllTransactions(ctx.Request().Context())
	if err != nil {
		logger.Error("failed to get all transactions", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get all transactions"})
	}

	return ctx.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) TransferFunds(ctx echo.Context) error {
	logger := h.Logger.GetLogger()

	var req models.TransferFundsRequest
	if err := ctx.Bind(&req); err != nil {
		logger.Error("failed to parse request", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.SenderID == uuid.Nil || req.ReceiverID == uuid.Nil || req.Amount <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid or missing parameters"})
	}

	accounts, err := h.TransactionService.TransferFunds(ctx.Request().Context(), req.SenderID, req.ReceiverID, req.Amount)
	if err != nil {
		logger.Error("failed to transfer funds", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to transfer funds"})
	}

	return ctx.JSON(http.StatusOK, accounts)

}
