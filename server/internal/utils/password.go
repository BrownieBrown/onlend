package utils

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCostFactor = 14

func GenerateHash(password string, costFactor int, l Logger) (string, error) {
	logger := l.GetLogger()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costFactor)
	if err != nil {
		logger.Error("Error generating hash", zap.Error(err))
		return "", err
	}
	return string(bytes), nil
}

func GenerateHashPassword(password string, l Logger) (string, error) {
	return GenerateHash(password, bcryptCostFactor, l)
}

func CompareHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func CompareHashPassword(password, hash string, l Logger) error {
	logger := l.GetLogger()
	err := CompareHash(password, hash)
	if err != nil {
		logger.Error("Error comparing password hash", zap.Error(err))
		return err
	}
	return nil
}
