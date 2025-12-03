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
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

type handler struct {
	JWKS *keyfunc.JWKS
}

type authRequest struct {
	IdToken string `json:"id_token"`
}

type authResponse struct {
	AccessToken string `json:"access_token"`
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

	if h.JWKS == nil {
		httpx.WriteJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	claims := &googleClaims{}

	token, parseErr := jwt.ParseWithClaims(authReq.IdToken, claims, h.JWKS.Keyfunc, jwt.WithValidMethods([]string{"RS256"}))

	if parseErr != nil || token == nil {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("ParseWithClaims returns error or token = nil")
		return
	}

	if _, methodOk := token.Method.(*jwt.SigningMethodRSA); !methodOk {
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
		log.Println("signing method does not belong to RSA family")
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

	tc := jwtx.AccessTokenClaims{
		Sub:   claims.Sub,
		Email: claims.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "personabox",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	signingKey := os.Getenv("PERSONABOX_SECRET")

	// GenerateAccessToken() method checks if signingKey != "". Returns signed token
	signedToken, tokenErr := jwtx.GenerateAccessToken(signingKey, tc)
	if tokenErr != nil {
		httpx.WriteJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(tokenErr.Error())
	}

	authRes := authResponse{
		AccessToken: signedToken,
	}

	httpx.WriteJSONSuccess(w, authRes)
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

	if claims.Sub == "" {
		return errors.New("sub not found")
	}

	if claims.ExpiresAt == nil || time.Now().After(claims.ExpiresAt.Time.Add(30*time.Second)) {
		return errors.New("token expired")
	}

	if claims.IssuedAt == nil || claims.IssuedAt.Time.After(time.Now().Add(30*time.Second)) {
		return errors.New("invalid issue date")
	}

	return nil
}
