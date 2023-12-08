package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
