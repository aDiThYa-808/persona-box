package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
	"github.com/google/uuid"
)

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

	ctx := r.Context()
	userID := claims.Sub
	personaID := uuid.New().String()

	persona := models.Persona{
		PersonaID:          personaID,
		UserID:             userID,
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
	}

	createErr := dynamodbx.CreateNewPersona(ctx, persona)
	if createErr != nil {
		httpx.WriteJSONError(w, "failed to create persona", http.StatusInternalServerError)
		return
	}

	httpx.WriteJSONSuccess(w, map[string]string{"message": "successfully created persona"})
}
