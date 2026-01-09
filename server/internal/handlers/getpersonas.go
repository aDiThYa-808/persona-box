package handlers

import (
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func GetPersonasHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx := r.Context()
	userID := claims.Sub

	personas, getErr := dynamodbx.GetUsersPersonas(ctx, userID)
	if getErr != nil {
		httpx.WriteJSONError(w, "failed to get personas", http.StatusInternalServerError)
		return
	}

	httpx.WriteJSONSuccess(w, personas)

}
