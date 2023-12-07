package utils

import (
	"errors"
	"os"
	"server/pkg/models"
)

// LoadConfig loads all configuration settings
func LoadConfig() (models.DatabaseConfig, error) {
	postgresConfig, err := loadPostgresConfig()
	if err != nil {
		return models.DatabaseConfig{}, err
	}

	// Load other configurations as necessary

	return models.DatabaseConfig{
		Postgres: postgresConfig,
		// Initialize other configurations here
	}, nil
}

// loadPostgresConfig loads PostgresSQL specific configuration
func loadPostgresConfig() (models.PostgresConfig, error) {
	config := models.PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_NAME"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	}

	if config.Host == "" || config.Port == "" || config.User == "" || config.Password == "" || config.DBName == "" || config.SSLMode == "" {
		return models.PostgresConfig{}, errors.New("missing required PostgresSQL environment variables")
	}

	return config, nil
}
