package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"server/internal/utils"
	"server/pkg/models"
	"time"
)

type UserService struct {
	Repository     models.UserRepository
	AccountService *AccountService
	Config         models.Config
}

func NewUserService(repository models.UserRepository, accountService *AccountService, cfg models.Config) *UserService {
	return &UserService{
		Repository:     repository,
		AccountService: accountService,
		Config:         cfg,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	response, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	defaultAccount := models.Account{
		Id:          uuid.New(),
		UserID:      response.Id,
		AccountType: "default",
		Balance:     0,
	}

	err = s.AccountService.CreateAccount(ctx, &defaultAccount)
	if err != nil {
		return nil, err
	}

	res := &models.CreateUserResponse{
		Id:       response.Id.String(),
		Username: response.Username,
		Email:    response.Email,
	}

	return res, nil
}

func (s *UserService) Login(ctx context.Context, req *models.LoginUserRequest) (*models.LoginUserResponse, error) {
	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CompareHashPassword(req.Password, user.Password)
	if err != nil {
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
		return nil, err
	}

	res := &models.LoginUserResponse{
		AccessToken: session,
		Id:          user.Id.String(),
		Username:    user.Username,
	}

	return res, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.GetUserResponse, error) {
	users, err := s.Repository.GetAllUsers(ctx)
	if err != nil {
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
