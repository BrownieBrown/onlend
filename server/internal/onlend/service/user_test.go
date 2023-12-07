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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockUserRepository(ctrl)
	l, err := utils.NewZapLogger()
	assert.NoError(t, err, "Unexpected error creating logger")

	timeout := time.Second * 5
	userService := service.NewUserService(mockRepository, l, timeout)

	ctx := context.Background()
	req := &models.CreateUserRequest{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password",
	}

	id := uuid.New()
	mockRepository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&models.User{
		Id:       id,
		Username: req.Username,
		Email:    req.Email,
		Password: "hashed_password",
	}, nil)

	response, err := userService.CreateUser(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, id.String(), response.Id)
	assert.Equal(t, req.Username, response.Username)
	assert.Equal(t, req.Email, response.Email)
}
