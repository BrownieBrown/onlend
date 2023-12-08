package rest_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"server/internal/onlend/rest"
	"server/internal/utils"
	"server/mocks"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"server/pkg/models"
)

func TestSuccessfulUserCreationWithValidData(t *testing.T) {
	defer utils.UnsetEnvVars()
	utils.SetEnvVars()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := utils.LoadConfig()
	assert.NoError(t, err, "Unexpected error loading config")

	mockUserService := mocks.NewMockUserService(ctrl)
	logger, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	handler := rest.NewUserHandler(mockUserService, logger, cfg)

	e := echo.New()
	userRequest := models.CreateUserRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password",
	}
	requestBody, _ := json.Marshal(userRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&models.CreateUserResponse{}, nil)

	err = handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestUserCreationReturnsNonEmptyID(t *testing.T) {
	defer utils.UnsetEnvVars()
	utils.SetEnvVars()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := utils.LoadConfig()
	assert.NoError(t, err, "Unexpected error loading config")

	mockUserService := mocks.NewMockUserService(ctrl)

	logger, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	handler := rest.NewUserHandler(mockUserService, logger, cfg)

	e := echo.New()
	userRequest := models.CreateUserRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password",
	}
	requestBody, _ := json.Marshal(userRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&models.CreateUserResponse{Id: "12345"}, nil)

	err = handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var res models.CreateUserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
}

func TestLogin(t *testing.T) {
	defer utils.UnsetEnvVars()
	utils.SetEnvVars()

	cfg, err := utils.LoadConfig()
	assert.NoError(t, err, "Unexpected error loading config")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	logger, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	handler := rest.NewUserHandler(mockUserService, logger, cfg)

	e := echo.New()
	email := "test@gmail.com"
	password := "password"
	userRequest := models.LoginUserRequest{
		Email:    email,
		Password: password,
	}
	requestBody, _ := json.Marshal(userRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	id := uuid.New().String()
	username := "testUser"
	res := &models.LoginUserResponse{
		Id:       id,
		Username: username,
	}

	mockUserService.EXPECT().Login(gomock.Any(), gomock.Any()).Return(res, nil)

	err = handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.LoginUserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
}
