package utils

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	l, err := NewZapLogger()
	if err != nil {
		return "", err
	}
	logger := l.GetLogger()

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Error("Error generating password hash", zap.Error(err))
		return "", err
	}
	return string(bytes), nil
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
