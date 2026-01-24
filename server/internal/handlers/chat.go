package handlers

import (
	"encoding/json"
	"log"
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
	Title     string `json:"title"`
	Response  string `json:"response"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	accessToken, extractErr := httpx.ExtractToken(r)
	if extractErr != nil {
		httpx.WriteJSONError(w, extractErr.Error(), http.StatusUnauthorized)
		return
	}

	claims, validateErr := jwtx.ValidateAccessToken(accessToken)
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

	persona, getPersonaErr := dynamodbx.GetPersonaByPersonaID(ctx, claims.Sub, req.PersonaID)
	if getPersonaErr != nil {
		log.Println(getPersonaErr)
		httpx.WriteJSONError(w, "persona doesnt exist", http.StatusNotFound)
		return
	}

	systemPrompt := openaiadapter.CreateSystemPrompt(persona)
	assistantMessage := ""
	userMessage := req.Message

	resp, tokensUsed, chatErr := openaiadapter.Chat(ctx, systemPrompt, assistantMessage, userMessage)
	if chatErr != nil {
		log.Println(chatErr)
		httpx.WriteJSONError(w, "failed to access llm", http.StatusInternalServerError)
		return
	}

	if req.SessionID == "" {
		title, summary, err := openaiadapter.GenerateTitleAndSummary(ctx, userMessage)
		if err != nil {
			log.Println(err)
			title = persona.PersonaName + " - New Chat"
			summary = userMessage + "." + resp
		}

		chatSession := models.ChatSession{
			SessionID:    uuid.New().String(),
			PersonaID:    req.PersonaID,
			Title:        title,
			CreatedAt:    time.Now().UTC().Format(time.RFC3339),
			UpdatedAt:    time.Now().UTC().Format(time.RFC3339),
			MessageCount: 2,
			TokenCount:   int(tokensUsed),
			Summary:      summary,
		}

		createErr := dynamodbx.CreateNewChatSession(ctx, chatSession)
		if createErr != nil {
			log.Println(createErr)
			httpx.WriteJSONError(w, "couldnt create chat session", http.StatusInternalServerError)
			return
		}

		response.SessionID = chatSession.SessionID
		response.Title = chatSession.Title
	} else {
		chatSession, getChatSessionErr := dynamodbx.GetChatSessionBySessionID(ctx, req.SessionID)
		if getChatSessionErr != nil {
			log.Println(getChatSessionErr)
			httpx.WriteJSONError(w, "chat session not found", http.StatusNotFound)
			return
		}

		if chatSession.TokenCount+int(tokensUsed) > 7000 { // add message limit after deciding the limit.
			httpx.WriteJSONError(w, "chat limit reached", http.StatusTooManyRequests)
			return
		}

		// update summary after every N number of messages. N tbd

		updateErr := dynamodbx.UpdateSessionMessageAndTokenCount(ctx, chatSession.SessionID, time.Now().UTC().Format(time.RFC3339), 2, int(tokensUsed))
		if updateErr != nil {
			log.Println(updateErr)
			httpx.WriteJSONError(w, "failed to update chat session", http.StatusInternalServerError)
			return
		}
	}

	// store both the user and assistant message in the messages tables

	response.Response = resp

	httpx.WriteJSONSuccess(w, response)
}
