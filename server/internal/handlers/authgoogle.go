package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"

	"github.com/aDiThYa-808/persona-box/internal/httpx"
)

type handler struct {
	JWKS *keyfunc.JWKS
}

type authRequest struct {
	IdToken string `json:"id_token"`
}

type googleClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	jwt.RegisteredClaims
}

// constructor for dependency injection
func New(jwks *keyfunc.JWKS) *handler {
	return &handler{JWKS: jwks}
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

	claims := &googleClaims{}

	token, parseErr := jwt.ParseWithClaims(authReq.IdToken, claims, h.JWKS.Keyfunc)

	if parseErr != nil {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("ParseWithClaims failed.")
		return
	}

	if _, methodOk := token.Method.(*jwt.SigningMethodRSA); !methodOk {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("Token not valid")
		return
	}

	verifyErr := verifyClaims(claims)
	if verifyErr != nil {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println(verifyErr.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Valid JWT Token\n"))
}

// helpers
func isAudienceValid(audience jwt.ClaimStrings) error {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	if slices.Contains(audience, clientId) {
		return nil
	}
	return errors.New("invalid audience")
}

func verifyClaims(claims *googleClaims) error {
	if !claims.EmailVerified {
		return errors.New("email not verified")
	}

	if audErr := isAudienceValid(claims.Audience); audErr != nil {
		return audErr
	}

	if claims.Issuer != "https://accounts.google.com" && claims.Issuer != "accounts.google.com" {
		return errors.New("invalid issuer")
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		return errors.New("token expired")
	}

	return nil
}
