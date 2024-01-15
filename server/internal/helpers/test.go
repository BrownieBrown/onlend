package helpers

import (
	"github.com/google/uuid"
	"server/internal/utils"
	"server/pkg/models"
)

func CreateUser(username, email, password string) *models.User {
	hashedPassword, _ := utils.GenerateHashPassword(password)
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

func CreateAccount(accountId uuid.UUID, userID uuid.UUID, accountType string, balance float64) *models.Account {
	return &models.Account{
		Id:          accountId,
		UserID:      userID,
		AccountType: accountType,
		Balance:     balance,
	}
}

func CreateTransaction(transactionId, senderId, receiverId uuid.UUID, amount float64, transactionType models.TransactionType, status models.Status) *models.Transaction {
	return &models.Transaction{
		Id:              transactionId,
		SenderID:        senderId,
		ReceiverID:      receiverId,
		Amount:          amount,
		TransactionType: transactionType,
		Status:          status,
	}
}

func CreateTransferFundsRequest(senderId, receiverId uuid.UUID, amount float64) *models.TransferFundsRequest {
	return &models.TransferFundsRequest{
		SenderID:   senderId,
		ReceiverID: receiverId,
		Amount:     amount,
	}
}

func CreateUpdateAccountRequest(accountId uuid.UUID, sum float64, transactionType models.TransactionType) *models.UpdateAccountRequest {
	return &models.UpdateAccountRequest{
		Id:              accountId.String(),
		TransactionType: transactionType,
		Sum:             sum,
	}
}
