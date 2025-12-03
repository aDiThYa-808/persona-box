/*
contains function to validate access token.
called inside handlers to validate incoming requests.
*/
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

// verifies the provided access token and returns parsed claims(type JWTClaims) or an error if the token is invalid.
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
