package jwtx

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func generateTestToken(t *testing.T, signingKey string, claims jwt.Claims) string {
	testToken, tokenErr := GenerateAccessToken(signingKey, claims)
	if tokenErr != nil {
		t.Fatalf("failed to generate test token: %v", tokenErr)
	}
	return testToken
}
