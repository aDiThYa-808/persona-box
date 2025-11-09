package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"

	"github.com/aDiThYa-808/persona-box/internal/httpx"
)

type handler struct {
	JWKS *keyfunc.JWKS
}

// constructor for dependency injection
func New(jwks *keyfunc.JWKS) *handler {
	return &handler{JWKS: jwks}
}

type authRequest struct {
	IdToken string `json:"id_token"`
}

func (h *handler) GoogleAuthHandler(w http.ResponseWriter, r *http.Request) {

	var authReq authRequest
	decodeReqErr := json.NewDecoder(r.Body).Decode(&authReq)

	if decodeReqErr != nil || authReq.IdToken == "" {
		httpx.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Could not decode request body.")
		return
	}

	defer r.Body.Close()

	claims := &jwt.RegisteredClaims{}

	token, parseErr := jwt.ParseWithClaims(authReq.IdToken, claims, h.JWKS.Keyfunc)

	if parseErr != nil {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("ParseWithClaims failed.")
		return
	}

	if !token.Valid {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("Token not valid")
		return
	}

	if !verifyClaims(claims) {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("Invalid Audience, Issuer or ExpiresAt")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Valid JWT Token\n"))
}

// helpers
func isAudienceValid(Audience jwt.ClaimStrings) bool {
	audValid := false
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	for _, a := range Audience {
		if a == clientId {
			audValid = true
			break
		}
	}
	return audValid
}

func verifyClaims(claims *jwt.RegisteredClaims) bool {
	if !isAudienceValid(claims.Audience) {
		return false
	}

	if claims.Issuer != "https://accounts.google.com" && claims.Issuer != "accounts.google.com" {
		return false
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return false
	}

	return true
}
