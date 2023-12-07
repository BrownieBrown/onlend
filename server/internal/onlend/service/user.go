package service

import (
	"context"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
	"time"
)

type service struct {
	Repository models.UserRepository
	timeout    time.Duration
}

func NewUserService(repository models.UserRepository) models.UserService {
	return &service{repository, time.Duration(2) * time.Second}
}

func (s *service) CreateUser(c context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	l, err := utils.NewZapLogger()
	if err != nil {
		return nil, err
	}
	logger := l.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.GenerateHashPassword(req.Password, l)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	response, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		logger.Error("Error creating user", zap.Error(err))
		return nil, err
	}

	res := &models.CreateUserResponse{
		Id:       response.Id.String(),
		Username: response.Username,
		Email:    response.Email,
	}

	return res, nil
}
