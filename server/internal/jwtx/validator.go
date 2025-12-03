/*
validator.go contains functions to validate tokens.
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

// Parses idToken and populates the claims struct.
func ParseGoogleJWT(idToken string, claims *GoogleClaims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
	token, parseErr := jwt.ParseWithClaims(idToken, claims, keyFunc, jwt.WithValidMethods([]string{"RS256"}))
	if parseErr != nil || token == nil {
		return nil, errors.New("couldnt parse id token")
	}
	if _, methodOk := token.Method.(*jwt.SigningMethodRSA); !methodOk {
		return nil, errors.New("invalid signing method")
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return token, nil
}
