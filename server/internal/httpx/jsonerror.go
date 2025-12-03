package httpx

import (
	"encoding/json"
	"net/http"
)

// returns http response with provided error message and status in JSON
func WriteJSONError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}
