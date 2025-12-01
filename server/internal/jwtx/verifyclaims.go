package jwtx

import (
	"errors"
	"time"
)

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
