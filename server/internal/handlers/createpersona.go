package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
	"github.com/google/uuid"
)

type createPersonaResponse struct {
	PersonaID   string `json:"persona_id"`
	PersonaName string `json:"name"`
	CreatedAt   string `json:"created_at"`
}

func CreatePersonaHandler(w http.ResponseWriter, r *http.Request) {
	token, extractErr := httpx.ExtractToken(r)
	if extractErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims, validateErr := jwtx.ValidateAccessToken(token)
	if validateErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req models.Persona
	decodeErr := json.NewDecoder(r.Body).Decode(&req)
	if decodeErr != nil {
		httpx.WriteJSONError(w, "bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	req.UserID = claims.Sub

	// assign unique persona id(sort key) for the persona
	req.PersonaID = uuid.New().String()

	persona := models.Persona{
		PersonaID:          req.PersonaID,
		UserID:             req.UserID,
		PersonaName:        req.PersonaName,
		PersonaDescription: req.PersonaDescription,
		Age:                req.Age,
		Pronouns:           req.Pronouns,
		Openness:           req.Openness,
		Conscientiousness:  req.Conscientiousness,
		Extraversion:       req.Extraversion,
		Agreeableness:      req.Agreeableness,
		Neuroticism:        req.Neuroticism,
		Intelligence:       req.Intelligence,
		ThinkingStyle:      req.ThinkingStyle,
		Tone:               req.Tone,
		HumorLevel:         req.HumorLevel,
		MoodFluctuation:    req.MoodFluctuation,
		Likes:              req.Likes,
		Dislikes:           req.Dislikes,
		Formality:          req.Formality,
		Fluency:            req.Fluency,
		EmojiUsage:         req.EmojiUsage,
		ResponseLength:     req.ResponseLength,
		CreatedAt:          time.Now().UTC().Format(time.RFC3339),
	}

	createErr := dynamodbx.CreateNewPersona(ctx, persona)
	if createErr != nil {
		httpx.WriteJSONError(w, "failed to create persona", http.StatusInternalServerError)
		return
	}

	response := createPersonaResponse{
		PersonaID:   persona.PersonaID,
		PersonaName: persona.PersonaName,
		CreatedAt:   persona.CreatedAt,
	}

	httpx.WriteJSONSuccess(w, response)
}
