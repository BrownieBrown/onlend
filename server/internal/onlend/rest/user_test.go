package rest_test

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"server/internal/helpers"
	"server/internal/onlend/rest"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSuccessfulUserCreationWithValidData(t *testing.T) {
	mockUserService, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()
	userRequest := helpers.CreateUserRequest("testUser", "test@gmail.com", "password")
	c, rec := prepareTestRequest(e, userRequest, "/api/v1/signup", "post")

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&models.CreateUserResponse{}, nil)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestUserCreationReturnsNonEmptyID(t *testing.T) {
	mockUserService, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()

	userRequest := helpers.CreateUserRequest("testUser", "test@gmail.com", "password")
	c, rec := prepareTestRequest(e, userRequest, "/api/v1/signup", "post")

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&models.CreateUserResponse{Id: "12345"}, nil)

	err := handler.CreateUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var res models.CreateUserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
}

func TestLogin(t *testing.T) {
	mockUserService, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()

	userRequest := helpers.CreateLoginUserReq("test@gmail.com", "password")
	c, rec := prepareTestRequest(e, userRequest, "/api/v1/login", "post")

	id := uuid.New().String()
	username := "testUser"
	res := helpers.CreateLoginUserResponse(id, username, "token")

	mockUserService.EXPECT().Login(gomock.Any(), gomock.Any()).Return(res, nil)

	err := handler.Login(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.LoginUserResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Id)
}

func TestGetAllUsers(t *testing.T) {
	mockUserService, handler, ctrl, _ := setupUserHandler(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	e := echo.New()
	c, rec := prepareTestRequest(e, nil, "/api/v1/users", "get")

	user1 := helpers.CreateGetUserResponse("1", "testUser", "test1@gmail.com")
	user2 := helpers.CreateGetUserResponse("1", "testUser2", "test2@gmail.com")
	users := []*models.GetUserResponse{user1, user2}
	mockUserService.EXPECT().GetAllUsers(gomock.Any()).Return(users, nil)

	err := handler.GetAllUsers(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func setupUserHandler(t *testing.T) (*mocks.MockUserService, *rest.UserHandler, *gomock.Controller, utils.Logger) {
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

	return mockUserService, handler, ctrl, logger
}

func prepareTestRequest(e *echo.Echo, requestBody interface{}, target string, method string) (echo.Context, *httptest.ResponseRecorder) {
	var reqBodyBuffer *bytes.Buffer
	if requestBody != nil {
		marshaledBody, _ := json.Marshal(requestBody)
		reqBodyBuffer = bytes.NewBuffer(marshaledBody)
	} else {
		reqBodyBuffer = bytes.NewBuffer(nil)
	}

	req := httptest.NewRequest(method, target, reqBodyBuffer)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}
