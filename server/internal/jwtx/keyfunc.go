package jwtx

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func keyFunc(t *jwt.Token) (interface{}, error) {
	_, ok := t.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("unexpected claims type")
	}

	if _, methodOk := t.Method.(*jwt.SigningMethodHMAC); !methodOk {
		return nil, errors.New("invalid signing method")
	}

	signingKey := os.Getenv("PERSONABOX_SECRET")

	if signingKey == "" {
		return nil, errors.New("signing key not found")
	}

	return []byte(signingKey), nil
}
