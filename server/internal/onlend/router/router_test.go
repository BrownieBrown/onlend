package router_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"server/internal/helpers"
	"server/internal/onlend/rest"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"strings"
	"testing"
)

func TestSuccessfulUserCreation(t *testing.T) {
	router, handler, ctrl, mockUserService := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "John Doe", "email": "johndoe@example.com", "password": "password"}`
	c, rec := prepareTestRequest(router, reqBody, "/api/v1/signup", http.MethodPost)

	user := helpers.CreateUserResponse(uuid.New().String(), "John Doe", "john@gmail.com")
	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(user, nil)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var res models.User

	err = json.Unmarshal(rec.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, res.Id.String())
	assert.Equal(t, user.Username, res.Username)
	assert.Equal(t, user.Email, res.Email)
}

func TestUserCreationWithInvalidEmail(t *testing.T) {
	router, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "John Doe", "email": "", "password": "password"}`
	c, rec := prepareTestRequest(router, reqBody, "/api/v1/signup", http.MethodPost)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUserCreationWithInvalidPassword(t *testing.T) {
	router, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "John Doe", "email": "john.doe@gmail.com", "password": ""}`
	c, rec := prepareTestRequest(router, reqBody, "/api/v1/signup", http.MethodPost)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUserCreationWithInvalidUsername(t *testing.T) {
	router, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "", "email": "john.doe@gmail.com", "password": "password"}`
	c, rec := prepareTestRequest(router, reqBody, "/api/v1/signup", http.MethodPost)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func setupUserHandler(t *testing.T) (*echo.Echo, *rest.UserHandler, *gomock.Controller, *mocks.MockUserService) {
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

	mockUserService := mocks.NewMockUserService(ctrl)

	handler := rest.NewUserHandler(mockUserService, logger, cfg)

	router := echo.New()
	return router, handler, ctrl, mockUserService
}

func prepareTestRequest(router *echo.Echo, requestBody, target, method string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, target, strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := router.NewContext(req, rec)

	return c, rec
}
