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

	accessToken, exists := strings.CutPrefix(header, "Bearer ")

	if !exists || accessToken == "" {
		return "", errors.New("invalid authorization header")
	}

	return accessToken, nil
}
