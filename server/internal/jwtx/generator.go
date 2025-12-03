/*
generator.go contains function to generate access token
*/

package jwtx

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// Uses the provided signing key and claims to create a JWT string. Returns error if the process failed.
func GenerateAccessToken(signingKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if signingKey == "" {
		return "", errors.New("missing signing key")
	}

	signedToken, tokenErr := token.SignedString([]byte(signingKey))

	if tokenErr != nil {
		return "", errors.New("error occured while creating signed jwt")
	}

	return signedToken, nil
}
