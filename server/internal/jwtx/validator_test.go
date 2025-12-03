package jwtx

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func init() {
	// use the real signing key for testing.
	loadErr := godotenv.Load("../../.env")
	if loadErr != nil {
		log.Fatal("couldnt load env")
	}
}

func TestValidateAccessToken(t *testing.T) {
	signingKey := os.Getenv("PERSONABOX_SECRET")

	validClaims := AccessTokenClaims{
		Sub:   "test-user",
		Email: "test@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "personabox",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	expiredClaims := AccessTokenClaims{
		Sub:   "test-user",
		Email: "test@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "personabox",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		},
	}

	invalidClaims := AccessTokenClaims{
		Sub:   "test-user",
		Email: "test@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "invalid-issuer",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{"valid token", generateTestToken(t, signingKey, validClaims), false},
		{"expired token", generateTestToken(t, signingKey, expiredClaims), true},
		{"invalid token", generateTestToken(t, signingKey, invalidClaims), true},
		{"malformed token", "not-a-token", true},
		{"empty token", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ValidateAccessToken(tc.token)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error=%v, got=%v", tc.wantErr, err)
			}
		})
	}
}
