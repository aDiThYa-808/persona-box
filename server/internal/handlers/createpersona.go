package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/dynamodbx/models"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
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
	personaID := "" // create unique id

	persona := models.Persona{
		UserID:    userID,
		PersonaID: personaID,
		//add the remaining fields
	}

	createErr := dynamodbx.CreateNewPersona(ctx, persona)
	if createErr != nil {
		httpx.WriteJSONError(w, "failed to create persona", http.StatusInternalServerError)
		return
	}
}
