package httpx

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("authorization header missing")
	}

	token, exists := strings.CutPrefix(header, "Bearer ")

	if !exists || token == "" {
		return "", errors.New("invalid authorization header")
	}

	return token, nil
}