package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
	"github.com/aDiThYa-808/persona-box/internal/openaiadapter"
	"github.com/google/uuid"
)

type ChatRequest struct {
	SessionID string `json:"session_id"`
	PersonaID string `json:"persona_id"`
	Message   string `json:"message"`
}

type ChatResponse struct {
	SessionID string `json:"session_id"`
	Response  string `json:"response"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	accessToken, extractErr := httpx.ExtractToken(r)
	if extractErr != nil {
		httpx.WriteJSONError(w, extractErr.Error(), http.StatusUnauthorized)
		return
	}

	_, validateErr := jwtx.ValidateAccessToken(accessToken)
	if validateErr != nil {
		httpx.WriteJSONError(w, validateErr.Error(), http.StatusUnauthorized)
		return
	}

	var req ChatRequest
	decodeErr := json.NewDecoder(r.Body).Decode(&req)

	if decodeErr != nil || req.Message == "" || req.PersonaID == "" {
		httpx.WriteJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	ctx := r.Context()

	response := ChatResponse{}

	//if client didnt send a session_id, it means they want to create a new chat session.
	if req.SessionID == "" {
		title:= " "
		summary := " "

		chatSession := models.ChatSession{
			SessionID:    uuid.New().String(),
			PersonaID:    req.PersonaID,
			Title:        title,
			CreatedAt:    time.Now().UTC().Format(time.RFC3339),
			UpdatedAt:    time.Now().UTC().Format(time.RFC3339),
			MessageCount: 0,
			TokenCount:   0,
			Summary:      summary,
		}

		createErr := dynamodbx.CreateNewChatSession(ctx, chatSession)
		if createErr != nil {
			httpx.WriteJSONError(w, "failed to create new chat", http.StatusInternalServerError)
			return
		}

		response.SessionID = chatSession.SessionID
	}

	systemMessage := ""
	assistantMessage := ""
	userMessage := req.Message

	responseMessage, _, chatErr := openaiadapter.Chat(ctx, systemMessage, assistantMessage, userMessage)
	if chatErr != nil {
		httpx.WriteJSONError(w, "failed to access llm", http.StatusInternalServerError)
		return
	}

	response.Response = responseMessage

	httpx.WriteJSONSuccess(w, response)
}
