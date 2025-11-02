package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aDiThYa-808/persona-box/internal/httpx"
)

type ChatRequest struct {
	Message string `json:"message"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.Message == "" {
		httpx.WriteJSONError(w, "Missing 'message' field in the request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	openaiKey := os.Getenv("OPENAI_SECRET_KEY")

	if openaiKey == "" {
		httpx.WriteJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Couldnt find openai secret key")
		return
	}
}
