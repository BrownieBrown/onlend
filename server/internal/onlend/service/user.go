package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"server/internal/utils"
	"server/pkg/models"
	"time"

	"github.com/pkg/errors"
)

type UserService struct {
	Repository     models.UserRepository
	AccountService *AccountService
	Logger         utils.Logger
	Timeout        time.Duration
	Config         models.Config
}

const (
	failedToHashPasswordErrMsg = "failed to hash password"
	errorCreatingUserErrMsg    = "Error creating user"
)

func NewUserService(repository models.UserRepository, accountService *AccountService, logger utils.Logger, timeout time.Duration, cfg models.Config) *UserService {
	return &UserService{
		Repository:     repository,
		AccountService: accountService,
		Logger:         logger,
		Timeout:        timeout,
		Config:         cfg,
	}
}

func (s *UserService) CreateUser(c context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
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

	defaultAccount := models.Account{
		Id:          uuid.New(),
		UserID:      response.Id,
		AccountType: "default",
		Balance:     0,
	}

	err = s.AccountService.CreateAccount(ctx, &defaultAccount)
	if err != nil {
		logger.Error("Error creating default account", zap.Error(err))
		return nil, err
	}

	res := &models.CreateUserResponse{
		Id:       response.Id.String(),
		Username: response.Username,
		Email:    response.Email,
	}

	return res, nil
}

func (s *UserService) Login(c context.Context, req *models.LoginUserRequest) (*models.LoginUserResponse, error) {
	logger := s.Logger.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Error("Error while finding user by email", zap.Error(err))
		return nil, err
	}

	err = utils.CompareHashPassword(req.Password, user.Password, s.Logger)
	if err != nil {
		logger.Error("Error while comparing password", zap.Error(err))
		return nil, err
	}

	jwtSigningMethod := jwt.SigningMethodHS256
	issuer := "Onlend"
	expirationTime := time.Duration(s.Config.JWT.JWTExpirationTime) * time.Second
	expires := jwt.NewNumericDate(time.Now().Add(expirationTime))
	jwtClaims := &models.JWTClaims{
		Id:       user.Id.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: expires,
		},
	}
	token := jwt.NewWithClaims(jwtSigningMethod, jwtClaims)
	session, err := token.SignedString([]byte(s.Config.JWT.JWTSigningKey))
	if err != nil {
		logger.Error("Error while signing JWT", zap.Error(err))
		return nil, err
	}

	res := &models.LoginUserResponse{
		AccessToken: session,
		Id:          user.Id.String(),
		Username:    user.Username,
	}

	return res, nil
}

func (s *UserService) GetAllUsers(c context.Context) ([]*models.GetUserResponse, error) {
	logger := s.Logger.GetLogger()

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	users, err := s.Repository.GetAllUsers(ctx)
	if err != nil {
		logger.Error("Error while getting all users", zap.Error(err))
		return nil, err
	}

	var res []*models.GetUserResponse
	for _, user := range users {
		u := models.GetUserResponse{
			Id:       user.Id.String(),
			Username: user.Username,
			Email:    user.Email,
		}
		res = append(res, &u)
	}

	return res, nil
}
