package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

type request struct {
	IdToken string `json:"id_token"`
}

func extractToken(r *http.Request) (string, error) {
	var req request
	decodeErr := json.NewDecoder(r.Body).Decode(&req)

	if decodeErr != nil {
		return "", errors.New("could not find id_token")
	}

	defer r.Body.Close()

	return req.IdToken, nil
}
