// GenerateHashPassword generates a hash from a password
package utils_test

import (
	"github.com/stretchr/testify/assert"
	"server/internal/utils"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCostFactor = 14

func TestGenerateHashPassword(t *testing.T) {
	l, err := utils.NewZapLogger()
	if err != nil {
		t.Errorf("Unexpected error creating logger: %v", err)
	}

	password := "password"
	hash, err := utils.GenerateHashPassword(password, l)
	assert.NoError(t, err, "Unexpected error generating hash")
	assert.NoError(t, err)
	assert.Greater(t, len(hash), 50, "Hashed password should be longer than 50 characters")

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.NoError(t, err, "Generated hash does not match password")
}

func TestCompareHashPassword(t *testing.T) {
	password := "password"
	wrongPassword := "wrongpassword"
	l, err := utils.NewZapLogger()
	if err != nil {
		t.Errorf("Unexpected error creating logger: %v", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCostFactor)
	assert.NoError(t, err, "Error generating hash for test")

	err = utils.CompareHashPassword(password, string(hash), l)
	assert.NoError(t, err, "Should correctly validate matching password")

	err = utils.CompareHashPassword(wrongPassword, string(hash), l)
	assert.Error(t, err, "Should fail to validate incorrect password")
}
