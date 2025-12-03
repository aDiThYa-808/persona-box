/*
verfiyclaims.go contains methods to verify provided claims
*/

package jwtx

import (
	"errors"
	"os"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// verifies the claims of personabox's access token
func verifyClaims(claims AccessTokenClaims) error {
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

// verifies the claims of Google's JWT
func VerifyGoogleClaims(claims *GoogleClaims) error {
	if !claims.EmailVerified {
		return errors.New("email not verified")
	}

	if audErr := isAudienceValid(claims.Audience); audErr != nil {
		return audErr
	}

	if claims.Issuer != "https://accounts.google.com" && claims.Issuer != "accounts.google.com" {
		return errors.New("invalid issuer")
	}

	if claims.Sub == "" {
		return errors.New("sub not found")
	}

	if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time.Add(30*time.Second)) {
		return errors.New("token expired")
	}

	if claims.IssuedAt == nil || claims.IssuedAt.Time.After(time.Now().Add(30*time.Second)) {
		return errors.New("invalid issue date")
	}

	return nil
}

func isAudienceValid(audience jwt.ClaimStrings) error {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	if slices.Contains(audience, clientId) {
		return nil
	}
	return errors.New("invalid audience")
}
