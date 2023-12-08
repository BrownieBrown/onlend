package models

type Config struct {
	Postgres PostgresConfig
	JWT      JWTConfig
	Cookie   CookieConfig
}
