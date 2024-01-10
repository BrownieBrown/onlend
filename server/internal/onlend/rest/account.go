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

func (h *AccountHandler) GetAccount(c echo.Context) error {
	logger := h.Logger.GetLogger()
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("failed to parse uuid", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to parse uuid"})
	}

	account, err := h.AccountService.GetAccount(c.Request().Context(), id)
	if err != nil {
		logger.Error("failed to get account", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get account"})
	}

	return c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) GetAllAccounts(c echo.Context) error {
	logger := h.Logger.GetLogger()

	accounts, err := h.AccountService.GetAllAccounts(c.Request().Context())
	if err != nil {
		logger.Error("failed to get all accounts", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get all accounts"})
	}

	return c.JSON(http.StatusOK, accounts)
}
