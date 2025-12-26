package httpx

import (
	"errors"
	"net/http"
)

// Extracts access token from request cookies and return it. Returns error if not found.
func ExtractToken(r *http.Request) (string, error) {
	token, err := r.Cookie("pb_access_token")

	if err != nil {
		return "", errors.New("token missing")
	}

	return token.Value, nil
}
