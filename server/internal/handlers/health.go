package handlers

import (
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	accessToken, extractErr := httpx.ExtractToken(r)
	if extractErr != nil {
		httpx.WriteJSONError(w, extractErr.Error(), http.StatusBadRequest)
		return
	}

	claims, validateErr := jwtx.ValidateAccessToken(accessToken)
	if validateErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("OK" + " " + claims.Email + "\n"))
}
