package jwtx

import "github.com/golang-jwt/jwt/v4"

type GoogleClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	EmailVerified bool   `json:"email_verified"`
	jwt.RegisteredClaims
}

type AccessTokenClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
