package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"server/internal/helpers"
	"server/internal/onlend/service"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"testing"
	"time"
)

func TestCreateUser_ValidInput_ReturnsResponse(t *testing.T) {
	userService, mockRepository, ctrl, _ := setup(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	req := &models.CreateUserRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password",
	}

	id := uuid.New()
	user := &models.User{
		Id:       id,
		Username: req.Username,
		Email:    req.Email,
		Password: "hashed_password",
	}
	mockRepository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(user, nil)

	response, err := userService.CreateUser(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, id.String(), response.Id)
	assert.Equal(t, req.Username, response.Username)
	assert.Equal(t, req.Email, response.Email)
}

func TestLogin_ValidInput_ReturnsResponse(t *testing.T) {
	userService, mockRepository, ctrl, logger := setup(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	email := "test@gmail.com"
	password := "password"
	username := "testUser"

	request := helpers.CreateLoginUserReq(email, password)
	user := helpers.CreateUser(logger, username, email, password)

	mockRepository.EXPECT().GetUserByEmail(gomock.Any(), email).Return(user, nil)

	response, err := userService.Login(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, username, response.Username)
}

func TestGetAllUsers(t *testing.T) {
	userService, mockRepository, ctrl, logger := setup(t)
	defer ctrl.Finish()
	defer utils.UnsetEnvVars()

	ctx := context.Background()

	user1 := helpers.CreateUser(logger, "testUser1", "test1@gmail.com", "password")
	user2 := helpers.CreateUser(logger, "testUser2", "test2@gmail.com", "password")

	users := []*models.User{user1, user2}

	mockRepository.EXPECT().GetAllUsers(gomock.Any()).Return(users, nil)
	resp, err := userService.GetAllUsers(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func setup(t *testing.T) (*service.Service, *mocks.MockUserRepository, *gomock.Controller, utils.Logger) {
	utils.SetEnvVars()
	ctrl := gomock.NewController(t)
	mockRepository := mocks.NewMockUserRepository(ctrl)

	l, err := utils.NewZapLogger()
	if err != nil {
		t.Fatalf("Failed to load logger: %v", err)
	}

	cfg, err := utils.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	timeout := time.Second * 5
	userService := service.NewUserService(mockRepository, l, timeout, cfg)
	return userService, mockRepository, ctrl, l
}
