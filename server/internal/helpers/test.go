package helpers

import (
	"github.com/google/uuid"
	"server/internal/utils"
	"server/pkg/models"
)

func CreateUser(logger utils.Logger, username, email, password string) *models.User {
	hashedPassword, _ := utils.GenerateHashPassword(password, logger)
	return &models.User{
		Id:       uuid.New(),
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
}

func CreateLoginUserReq(email, password string) *models.LoginUserRequest {
	return &models.LoginUserRequest{
		Email:    email,
		Password: password,
	}
}

func CreateUserRequest(username, email, password string) *models.CreateUserRequest {
	return &models.CreateUserRequest{
		Username: username,
		Email:    email,
		Password: password,
	}
}

func CreateLoginUserResponse(id, username, token string) *models.LoginUserResponse {
	return &models.LoginUserResponse{
		AccessToken: token,
		Id:          id,
		Username:    username,
	}
}

func CreateGetUserResponse(id, username, email string) *models.GetUserResponse {
	return &models.GetUserResponse{
		Id:       id,
		Username: username,
		Email:    email,
	}
}

func CreateUserResponse(is, username, email string) *models.CreateUserResponse {
	return &models.CreateUserResponse{
		Id:       is,
		Username: username,
		Email:    email,
	}
}
