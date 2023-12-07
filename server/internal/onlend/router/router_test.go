package router_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"server/internal/onlend/rest"
	"server/mocks"
	"server/pkg/models"
	"strings"
	"testing"
)

func TestSuccessfulUserCreation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	handler := &rest.UserHandler{
		UserService: mockUserService,
	}

	router := echo.New()
	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "John Doe", "email": "johndoe@example.com", "password": "password"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := router.NewContext(req, rec)

	user := &models.CreateUserResponse{
		Id:       uuid.New().String(),
		Username: "John Doe",
		Email:    "john@gmail.com",
	}
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	handler := &rest.UserHandler{
		UserService: mockUserService,
	}

	router := echo.New()
	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "John Doe", "email": "", "password": "password"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := router.NewContext(req, rec)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUserCreationWithInvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)

	handler := &rest.UserHandler{
		UserService: mockUserService,
	}

	router := echo.New()
	router.POST("/api/v1/signup", handler.CreateUser)

	reqBody := `{"username": "John Doe", "email": "john.doe@gmail.com", "password": ""}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/signup", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := router.NewContext(req, rec)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
