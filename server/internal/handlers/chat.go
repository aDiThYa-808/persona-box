package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aDiThYa-808/persona-box/internal/auth"
	"github.com/aDiThYa-808/persona-box/internal/httpx"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type ChatRequest struct {
	Message string `json:"message"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {

	_, validateErr := auth.ValidateRequest(r)
	if validateErr != nil {
		httpx.WriteJSONError(w, validateErr.Error(), http.StatusUnauthorized)
	}

	var req ChatRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&req)

	if decodeErr != nil || req.Message == "" {
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

	client := openai.NewClient(
		option.WithAPIKey(openaiKey),
	)

	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You are a friendly and concise AI assistant. Keep responses short and clear."),
			openai.UserMessage(req.Message),
		},
		Model: openai.ChatModelGPT4_1Mini,
	}

	chatCompletion, chatErr := client.Chat.Completions.New(context.TODO(), params)

	if chatErr != nil {
		httpx.WriteJSONError(w, "Could not connect to the AI model", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"response": chatCompletion.Choices[0].Message.Content,
	})

}
