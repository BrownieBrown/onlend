package utils

import "os"

func UnsetEnvVars() {
	os.Unsetenv("POSTGRES_HOST")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_NAME")
	os.Unsetenv("POSTGRES_SSL_MODE")
	os.Unsetenv("JWT_EXPIRATION_TIME")
	os.Unsetenv("JWT_SECRET_KEY")
	os.Unsetenv("COOKIE_NAME")
	os.Unsetenv("COOKIE_PATH")
	os.Unsetenv("COOKIE_DOMAIN")
	os.Unsetenv("COOKIE_MAX_AGE")
	os.Unsetenv("COOKIE_SECURE")
	os.Unsetenv("COOKIE_HTTP_ONLY")
}
