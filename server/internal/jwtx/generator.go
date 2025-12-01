package jwtx

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

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
