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
	Logger      utils.Logger
	Config      models.Config
}

func NewUserHandler(us models.UserService, logger utils.Logger, cfg models.Config) *UserHandler {
	return &UserHandler{
		UserService: us,
		Logger:      logger,
		Config:      cfg,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	logger := h.Logger.GetLogger()
	var user models.CreateUserRequest

	if err := c.Bind(&user); err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to bind request data"})
	}

	if ok, err := utils.ValidateUserInput(&user); !ok {
		logger.Error("failed to validate user input", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	res, err := h.UserService.CreateUser(c.Request().Context(), &user)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	response := models.CreateUserResponse{
		Id:       res.Id,
		Username: res.Username,
		Email:    res.Email,
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c echo.Context) error {
	logger := h.Logger.GetLogger()
	var user models.LoginUserRequest

	if err := c.Bind(&user); err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to bind request data"})
	}

	u, err := h.UserService.Login(c.Request().Context(), &user)
	if err != nil {
		logger.Error("failed to login user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to login user"})
	}

	cookie := createCookie(h.Config.Cookie, u.AccessToken)
	c.SetCookie(cookie)
	res := &models.BasicUserResponse{
		Id:       u.Id,
		Username: u.Username,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Logout(c echo.Context) error {
	cookie := unsetCookie(h.Config.Cookie)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "logout successful"})
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	logger := h.Logger.GetLogger()

	users, err := h.UserService.GetAllUsers(c.Request().Context())
	if err != nil {
		logger.Error("failed to get all users", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get all users"})
	}

	return c.JSON(http.StatusOK, users)
}

func createCookie(cfg models.CookieConfig, accessToken string) *http.Cookie {
	return &http.Cookie{
		Name:     cfg.Name,
		Value:    accessToken,
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		MaxAge:   cfg.MaxAge,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HttpOnly,
	}
}

func unsetCookie(cfg models.CookieConfig) *http.Cookie {
	return &http.Cookie{
		Name:     cfg.Name,
		Value:    "",
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		MaxAge:   -1,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HttpOnly,
	}
}
