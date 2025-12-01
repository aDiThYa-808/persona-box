package jwtx

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	jwt.RegisteredClaims
}

func ValidateAccessToken(accessToken string) (JWTClaims, error) {

	claims := &JWTClaims{}
	token, parseErr := jwt.ParseWithClaims(accessToken, claims, keyFunc, jwt.WithValidMethods([]string{"HS256"}))

	if parseErr != nil || token == nil {
		return JWTClaims{}, errors.New("unauthorized")
	}

	if !token.Valid {
		return JWTClaims{}, errors.New("unauthorized")
	}

	verifyErr := verifyClaims(*claims)

	if verifyErr != nil {

		return JWTClaims{}, verifyErr
	}

	return *claims, nil
}
