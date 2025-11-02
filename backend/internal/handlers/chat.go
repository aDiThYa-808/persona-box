package handlers

import (
	"encoding/json"
	"net/http"
)

type ChatRequest struct {
	Message string `json:"message"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Missing 'message' field in the request body",
		})
	}

	defer r.Body.Close()
}
