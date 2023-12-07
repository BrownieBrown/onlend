package rest_test

import (
	"bytes"
	"encoding/json"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	logger, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	handler := rest.NewUserHandler(mockUserService, logger)

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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	logger, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	handler := rest.NewUserHandler(mockUserService, logger)

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
