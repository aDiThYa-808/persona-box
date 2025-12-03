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

func TestGenerateAccessToken(t *testing.T) {
	signingKey := os.Getenv("PERSONABOX_SECRET")
	if signingKey == "" {
		t.Fatalf("signing key not found")
	}

	validClaims := AccessTokenClaims{
		Sub:   "test-user",
		Email: "test@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "personabox",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	tests := []struct {
		name    string
		key     string
		claims  jwt.Claims
		wantErr bool
	}{
		{"valid signing key", signingKey, validClaims, false},
		{"no signing key", "", validClaims, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GenerateAccessToken(tc.key, tc.claims)
			if (err != nil) != tc.wantErr {
				t.Errorf("expcected error=%v, got=%v", tc.wantErr, err)
			}
		})
	}
}
