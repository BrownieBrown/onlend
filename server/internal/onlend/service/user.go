package service

import (
	"context"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
	"time"

	"github.com/pkg/errors"
)

type Service struct {
	Repository models.UserRepository
	Logger     utils.Logger
	Timeout    time.Duration
}

const (
	failedToHashPasswordErrMsg = "failed to hash password"
	errorCreatingUserErrMsg    = "Error creating user"
)

func NewUserService(repository models.UserRepository, logger utils.Logger, timeout time.Duration) *Service {
	return &Service{Repository: repository, Logger: logger, Timeout: timeout}
}

func (s *Service) CreateUser(c context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	logger := s.Logger.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	hashedPassword, err := utils.GenerateHashPassword(req.Password, s.Logger)
	if err != nil {
		logger.Error(failedToHashPasswordErrMsg, zap.Error(err))
		return nil, errors.Wrap(err, failedToHashPasswordErrMsg)
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	response, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		logger.Error(errorCreatingUserErrMsg, zap.Error(err))
		return nil, errors.Wrap(err, errorCreatingUserErrMsg)
	}

	res := &models.CreateUserResponse{
		Id:       response.Id.String(),
		Username: response.Username,
		Email:    response.Email,
	}

	return res, nil
}
