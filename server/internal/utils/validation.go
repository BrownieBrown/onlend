package utils

import (
	"errors"
	"net/mail"
	"server/pkg/models"
)

func ValidateUserInput(user *models.CreateUserRequest) (bool, error) {
	if ok, err := validateEmail(user.Email); !ok {
		return false, err
	}
	if ok, err := validateUsername(user.Username); !ok {
		return false, err
	}
	if ok, err := validatePassword(user.Password); !ok {
		return false, err
	}
	return true, nil
}

func validateEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is required")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, errors.New("invalid email address")
	}
	return true, nil
}

func validateUsername(username string) (bool, error) {
	if username == "" {
		return false, errors.New("username is required")
	}
	return true, nil
}

func validatePassword(password string) (bool, error) {
	if password == "" {
		return false, errors.New("password is required")
	}
	return true, nil
}
