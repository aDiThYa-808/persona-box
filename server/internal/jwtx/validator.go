/*
contains function to validate access token.
called inside handlers to validate incoming requests.
*/
package jwtx

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// verifies the provided access token and returns parsed claims(type JWTClaims) or an error if the token is invalid.
func ValidateAccessToken(accessToken string) (AccessTokenClaims, error) {

	claims := &AccessTokenClaims{}
	token, parseErr := jwt.ParseWithClaims(accessToken, claims, keyFunc, jwt.WithValidMethods([]string{"HS256"}))

	if parseErr != nil || token == nil {
		return AccessTokenClaims{}, errors.New("unauthorized")
	}

	if !token.Valid {
		return AccessTokenClaims{}, errors.New("unauthorized")
	}

	verifyErr := verifyClaims(*claims)

	if verifyErr != nil {

		return AccessTokenClaims{}, verifyErr
	}

	return *claims, nil
}
