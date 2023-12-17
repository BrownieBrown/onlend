package utils_test

import (
	"github.com/stretchr/testify/assert"
	"server/internal/helpers"
	"server/internal/utils"
	"testing"
)

func TestValidateUserInputWithValidInput(t *testing.T) {
	user := helpers.CreateUserRequest("John Doe", "john.doe@gmail.com", "password")
	ok, err := utils.ValidateUserInput(user)
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestValidateUserInputWithInvalidEmail(t *testing.T) {
	user := helpers.CreateUserRequest("John Doe", "", "password")
	ok, err := utils.ValidateUserInput(user)
	assert.Error(t, err)
	assert.False(t, ok)
}

func TestValidateUserInputWithInvalidUsername(t *testing.T) {
	user := helpers.CreateUserRequest("", "john.doe@gmail.com", "password")
	ok, err := utils.ValidateUserInput(user)
	assert.Error(t, err)
	assert.False(t, ok)
}

func TestValidateUserInputWithInvalidPassword(t *testing.T) {
	user := helpers.CreateUserRequest("John Doe", "john.doe@gmail.com", "")
	ok, err := utils.ValidateUserInput(user)
	assert.Error(t, err)
	assert.False(t, ok)
}
