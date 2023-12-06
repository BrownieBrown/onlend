package rest

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"server/internal/utils"
	"server/pkg/models"
)

type UserHandler struct {
	UserService models.UserService
}

func NewUserHandler(us models.UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	logger := utils.GetLogger()
	var u models.CreateUserRequest

	if err := c.Bind(&u); err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request data"})
	}

	res, err := h.UserService.CreateUser(c.Request().Context(), &u)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	return c.JSON(http.StatusCreated, res)
}
