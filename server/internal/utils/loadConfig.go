package utils

import (
	"errors"
	"os"
	"server/pkg/models"
	"strconv"
)

// LoadConfig loads all configuration settings
func LoadConfig() (models.Config, error) {
	postgresConfig, err := loadPostgresConfig()
	if err != nil {
		return models.Config{}, err
	}

	jwtConfig, err := loadJWTConfig()
	if err != nil {
		return models.Config{}, err
	}

	cookieConfig, err := loadCookieConfig()
	if err != nil {
		return models.Config{}, err
	}

	// Load other configurations as necessary

	return models.Config{
		Postgres: postgresConfig,
		JWT:      jwtConfig,
		Cookie:   cookieConfig,
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

func loadJWTConfig() (models.JWTConfig, error) {
	expirationStr := os.Getenv("JWT_EXPIRATION_TIME")
	expirationTime, err := strconv.ParseInt(expirationStr, 10, 64)
	if expirationStr == "" || err != nil {
		expirationTime = 3600 // Default value if not set or parsing error
	}

	jwtSigningKey := os.Getenv("JWT_SECRET_KEY")
	config := models.JWTConfig{
		JWTExpirationTime: expirationTime,
		JWTSigningKey:     jwtSigningKey,
	}

	if config.JWTSigningKey == "" || config.JWTExpirationTime == 0 {
		return models.JWTConfig{}, errors.New("missing required JWT environment variables")
	}

	return config, nil
}

func loadCookieConfig() (models.CookieConfig, error) {
	name := os.Getenv("COOKIE_NAME")
	path := os.Getenv("COOKIE_PATH")
	domain := os.Getenv("COOKIE_DOMAIN")
	maxAgeStr := os.Getenv("COOKIE_MAX_AGE")
	maxAge, err := strconv.Atoi(maxAgeStr)
	if err != nil {
		maxAge = 3600 // Default value if not set or parsing error
	}
	secure := os.Getenv("COOKIE_SECURE")
	httpOnly := os.Getenv("COOKIE_HTTP_ONLY")

	config := models.CookieConfig{
		Name:     name,
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure == "false",
		HttpOnly: httpOnly == "true",
	}

	if config.Name == "" || config.Path == "" || config.Domain == "" {
		return models.CookieConfig{}, errors.New("missing required Cookie environment variables")
	}

	return config, nil
}
