package handlers

import (
	"context"
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
	CreatedAt string `json:"created_at"`
}

type ChatResponse struct {
	SessionID string `json:"session_id"`
	Title     string `json:"title"`
	Response  string `json:"response"`
	Timestamp string `json:"timestamp"`
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

	ctx, cancel := context.WithTimeout(r.Context(), 180*time.Second)
	defer cancel()

	response := ChatResponse{}

	persona, getPersonaErr := dynamodbx.GetPersonaByPersonaID(ctx, claims.Sub, req.PersonaID)
	if getPersonaErr != nil {
		log.Println(getPersonaErr)
		httpx.WriteJSONError(w, "persona doesnt exist", http.StatusNotFound)
		return
	}

	systemPrompt := openaiadapter.CreateSystemPrompt(persona)
	userMessage := req.Message
	assistantMessage := openaiadapter.CreateAssistantPrompt(userMessage, userMessage) // FIX THIS ASAP. USE LAST N MESSAGES AND SUMMARY FROM TABLE

	resp, tokensUsed, chatErr := openaiadapter.Chat(ctx, systemPrompt, assistantMessage, userMessage)
	if chatErr != nil {
		log.Println(chatErr)
		httpx.WriteJSONError(w, "failed to access llm", http.StatusInternalServerError)
		return
	}

	sessionID := req.SessionID

	if sessionID == "" {
		title, summary, err := openaiadapter.GenerateTitleAndSummary(ctx, userMessage)
		if err != nil {
			log.Println(err)
			title = persona.PersonaName + " - New Chat"
			summary = userMessage + "." + resp
		}

		sessionID = uuid.New().String()

		chatSession := models.ChatSession{
			SessionID:    sessionID,
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
		chatSession, getChatSessionErr := dynamodbx.GetChatSessionByID(ctx, req.PersonaID, req.SessionID)
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

		updateErr := dynamodbx.UpdateSessionMessageAndTokenCount(ctx, chatSession.PersonaID, chatSession.SessionID, time.Now().UTC().Format(time.RFC3339), 2, int(tokensUsed))
		if updateErr != nil {
			log.Println(updateErr)
			httpx.WriteJSONError(w, "failed to update chat session", http.StatusInternalServerError)
			return
		}
	}

	um := models.ChatMessage{
		SessionID: sessionID,
		CreatedAt: req.CreatedAt,
		Role:      "user",
		Message:   userMessage,
	}
	am := models.ChatMessage{
		SessionID: sessionID,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		Role:      "assistant",
		Message:   assistantMessage,
	}
	userMessageErr := dynamodbx.StoreChatMessage(ctx, um)
	assistantMessageErr := dynamodbx.StoreChatMessage(ctx, am)

	if userMessageErr != nil {
		log.Println(userMessageErr)
		httpx.WriteJSONError(w, "failed to store message", http.StatusInternalServerError)
		return
	}

	if assistantMessageErr != nil {
		log.Println(assistantMessageErr)
		httpx.WriteJSONError(w, "failed to store message", http.StatusInternalServerError)
		return
	}

	response.Timestamp = am.CreatedAt

	response.Response = resp

	httpx.WriteJSONSuccess(w, response)
}
