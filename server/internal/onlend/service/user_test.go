package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"server/internal/onlend/service"
	"server/internal/utils"
	"server/mocks"
	"server/pkg/models"
	"testing"
	"time"
)

func TestCreateUser_ValidInput_ReturnsResponse(t *testing.T) {
	defer utils.UnsetEnvVars()
	utils.SetEnvVars()

	cfg, err := utils.LoadConfig()
	assert.NoError(t, err, "Unexpected error loading config")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockUserRepository(ctrl)
	l, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	timeout := time.Second * 5
	userService := service.NewUserService(mockRepository, l, timeout, cfg)

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
	defer utils.UnsetEnvVars()
	utils.SetEnvVars()

	cfg, err := utils.LoadConfig()
	assert.NoError(t, err, "Unexpected error loading config")

	// Test case: All environment variables are set
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockUserRepository(ctrl)
	l, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	timeout := time.Second * 5
	userService := service.NewUserService(mockRepository, l, timeout, cfg)

	ctx := context.Background()
	email := "test@gmail.com"
	password := "password"
	hashedPassword, err := utils.GenerateHashPassword(password, l)
	assert.NoError(t, err)

	id := uuid.New()
	username := "testuser"

	request := &models.LoginUserRequest{
		Email:    email,
		Password: password,
	}
	user := &models.User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	mockRepository.EXPECT().GetUserByEmail(gomock.Any(), email).Return(user, nil)

	response, err := userService.Login(ctx, request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, id.String(), response.Id)
	assert.Equal(t, username, response.Username)
}
