package jwtx

import "github.com/golang-jwt/jwt/v4"

type AccessTokenClaims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
