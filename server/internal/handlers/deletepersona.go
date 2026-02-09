package handlers

import (
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func DeletePersonaHandler(w http.ResponseWriter, r *http.Request) {
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
	personaid := r.PathValue("id")
	if personaid == "" {
		httpx.WriteJSONError(w, "missing personaid", http.StatusBadRequest)
		return
	}

	deleteErr := dynamodbx.DeletePersona(r.Context(), claims.Sub, personaid)
	if deleteErr != nil {
		httpx.WriteJSONError(w, "couldnt complete persona deletion", http.StatusInternalServerError)
		return
	}

	httpx.WriteJSONSuccess(w, "successfully deleted persona")
}
