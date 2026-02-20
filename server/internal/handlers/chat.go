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

	var (
		sessionID    string
		chatSession  models.ChatSession
		isNewSession bool
		messages     []models.ChatMessage
	)

	if req.SessionID == "" {
		isNewSession = true
		sessionID = uuid.New().String()
	} else {
		isNewSession = false
		sessionID = req.SessionID

		cs, getChatSessionErr := dynamodbx.GetChatSessionByID(ctx, req.PersonaID, sessionID)
		if getChatSessionErr != nil {
			httpx.WriteJSONError(w, "chat session not found", http.StatusNotFound)
			return
		}

		chatSession = cs

		msgs, getMessageErr := dynamodbx.GetAllSessionMessages(ctx, sessionID)
		if getMessageErr != nil {
			httpx.WriteJSONError(w, "couldnt fetch messages", http.StatusInternalServerError)
			return
		}
		messages = msgs
	}

	resp, tokensUsed, chatErr := openaiadapter.Chat(ctx, systemPrompt, messages, userMessage)
	if chatErr != nil {
		log.Println(chatErr)
		httpx.WriteJSONError(w, "failed to access llm", http.StatusInternalServerError)
		return
	}

	messages = append(messages, models.ChatMessage{SessionID: sessionID, CreatedAt: time.Now().UTC().Format(time.RFC3339), Role: "user", Message: userMessage})
	messages = append(messages, models.ChatMessage{SessionID: sessionID, CreatedAt: time.Now().UTC().Format(time.RFC3339), Role: "assistant", Message: resp})

	if isNewSession {

		title, summary, err := openaiadapter.GenerateTitleAndSummary(ctx, messages)
		if err != nil {
			log.Println(err)
			title = persona.PersonaName + " - New Chat"
			summary = "User: " + userMessage + "." + " Assistant: " + resp
		}

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
		if chatSession.TokenCount+int(tokensUsed) > 7000 { // add message limit after deciding the limit.
			httpx.WriteJSONError(w, "chat limit reached", http.StatusTooManyRequests)
			return
		}
		if chatSession.MessageCount%10 == 0 {
			title, summary, err := openaiadapter.GenerateTitleAndSummary(ctx, messages)
			if err == nil {
				dynamodbx.UpdateTitleAndSummary(ctx, chatSession.PersonaID, chatSession.SessionID, title, summary) // fail silently if update fail. previous summary remains
			}
		}

		updateErr := dynamodbx.UpdateSessionMessageAndTokenCount(ctx, chatSession.PersonaID, chatSession.SessionID, time.Now().UTC().Format(time.RFC3339), 2, int(tokensUsed))
		if updateErr != nil {
			log.Println(updateErr)
			httpx.WriteJSONError(w, "failed to update chat session", http.StatusInternalServerError)
			return
		}
	}

	// store user message
	um := models.ChatMessage{
		SessionID: sessionID,
		CreatedAt: req.CreatedAt,
		Role:      "user",
		Message:   userMessage,
	}
	userMessageErr := dynamodbx.StoreChatMessage(ctx, um)
	if userMessageErr != nil {
		log.Println(userMessageErr)
		httpx.WriteJSONError(w, "failed to store message", http.StatusInternalServerError)
		return
	}

	// store assistant response
	am := models.ChatMessage{
		SessionID: sessionID,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		Role:      "assistant",
		Message:   resp,
	}
	assistantMessageErr := dynamodbx.StoreChatMessage(ctx, am)
	if assistantMessageErr != nil {
		log.Println(assistantMessageErr)
		httpx.WriteJSONError(w, "failed to store message", http.StatusInternalServerError)
		return
	}

	response.Timestamp = am.CreatedAt

	response.Response = resp

	httpx.WriteJSONSuccess(w, response)
}
