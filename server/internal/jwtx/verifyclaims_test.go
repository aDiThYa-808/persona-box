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
	if loadErr := godotenv.Load("../../.env"); loadErr != nil {
		log.Fatal("couldnt load env")
	}
}

func TestVerifyGoogleClaims(t *testing.T) {
	aud := os.Getenv("GOOGLE_CLIENT_ID")

	validClaims := GoogleClaims{
		Sub:           "test-sub",
		Email:         "testemail@gmail.com",
		EmailVerified: true,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{aud},
			Issuer:    "https://accounts.google.com",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	invalidClaims := GoogleClaims{
		Sub:   "test-user",
		Email: "test@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"invalid-aud"},
			Issuer:    "invalid-issuer",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		},
	}

	tests := []struct {
		name    string
		claims  *GoogleClaims
		wantErr bool
	}{
		{"valid claims", &validClaims, false},
		{"invalid claims", &invalidClaims, true},
		{"missing claims", &GoogleClaims{}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := VerifyGoogleClaims(tc.claims)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error=%v, got=%v", tc.wantErr, err)
			}
		})
	}
}
