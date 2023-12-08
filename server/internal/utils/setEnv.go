package utils

import "os"

func SetEnvVars() {
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "user")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_NAME", "dbname")
	os.Setenv("POSTGRES_SSL_MODE", "disable")
	os.Setenv("JWT_EXPIRATION_TIME", "3600")
	os.Setenv("JWT_SECRET_KEY", "secret")
	os.Setenv("COOKIE_NAME", "jwt")
	os.Setenv("COOKIE_PATH", "/")
	os.Setenv("COOKIE_DOMAIN", "localhost")
	os.Setenv("COOKIE_MAX_AGE", "3600")
	os.Setenv("COOKIE_SECURE", "false")
	os.Setenv("COOKIE_HTTP_ONLY", "true")
}
