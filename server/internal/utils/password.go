package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCostFactor = 14

func GenerateHash(password string, costFactor int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costFactor)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GenerateHashPassword(password string) (string, error) {
	return GenerateHash(password, bcryptCostFactor)
}

func CompareHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func CompareHashPassword(password, hash string) error {
	err := CompareHash(password, hash)
	if err != nil {
		return err
	}
	return nil
}
