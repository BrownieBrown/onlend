package models

type JWTConfig struct {
	JWTExpirationTime int64  `env:"JWT_EXPIRATION_TIME" env-default:"15"`
	JWTSigningKey     string `env:"JWT_SECRET_KEY"`
}
