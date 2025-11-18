package handlers

import (
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/auth"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	claims, validateErr := auth.ValidateRequest(r)
	if validateErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("OK" + " " + claims.Email + "\n"))
}
