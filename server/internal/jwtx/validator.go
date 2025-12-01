package jwtx

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	jwt.RegisteredClaims
}

func ValidateRequest(idToken string) (JWTClaims, error) {

	claims := &JWTClaims{}
	token, parseErr := jwt.ParseWithClaims(idToken, claims, keyFunc, jwt.WithValidMethods([]string{"HS256"}))

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

func verifyClaims(claims JWTClaims) error {
	if claims.Sub == "" {
		return errors.New("unauthorized")
	}

	if claims.Email == "" {
		return errors.New("unauthorized")
	}

	if claims.Issuer != "personabox" {
		return errors.New("unauthorized")
	}

	if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time.Add(30*time.Second)) {
		return errors.New("unauthorized")
	}

	if claims.IssuedAt == nil || claims.IssuedAt.Time.After(time.Now().Add(30*time.Second)) {
		return errors.New("unauthorized")
	}

	return nil
}
