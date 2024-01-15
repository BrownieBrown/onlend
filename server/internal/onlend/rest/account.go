package rest

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"server/internal/utils"
	"server/pkg/models"
)

type AccountHandler struct {
	AccountService models.AccountService
	Logger         utils.Logger
	Config         models.Config
}

func NewAccountHandler(as models.AccountService, logger utils.Logger, cfg models.Config) *AccountHandler {
	return &AccountHandler{
		AccountService: as,
		Logger:         logger,
		Config:         cfg,
	}
}

func (h *AccountHandler) GetAccount(ctx echo.Context) error {
	logger := h.Logger.GetLogger()
	idStr := ctx.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	account, err := h.AccountService.GetAccount(ctx.Request().Context(), id)
	if err != nil {
		logger.Error("failed to get account", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get account"})
	}

	return ctx.JSON(http.StatusOK, account)
}

func (h *AccountHandler) GetAccountByUserId(ctx echo.Context) error {
	logger := h.Logger.GetLogger()
	idStr := ctx.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	account, err := h.AccountService.GetAccountByUserId(ctx.Request().Context(), id)
	if err != nil {
		logger.Error("failed to get account", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get account"})
	}

	return ctx.JSON(http.StatusOK, account)
}

func (h *AccountHandler) GetAllAccounts(ctx echo.Context) error {
	logger := h.Logger.GetLogger()

	accounts, err := h.AccountService.GetAllAccounts(ctx.Request().Context())
	if err != nil {
		logger.Error("failed to get all accounts", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get all accounts"})
	}

	return ctx.JSON(http.StatusOK, accounts)
}

func (h *AccountHandler) UpdateAccount(ctx echo.Context) error {
	logger := h.Logger.GetLogger()

	var req models.UpdateAccountRequest
	if err := ctx.Bind(&req); err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to bind request"})
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	account, err := h.AccountService.UpdateAccount(ctx.Request().Context(), id, req.Sum, req.TransactionType)
	if err != nil {
		logger.Error("failed to update account", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update account"})
	}

	return ctx.JSON(http.StatusOK, account)
}
