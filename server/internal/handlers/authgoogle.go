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

	claims := &jwtx.GoogleClaims{}

	_, parseErr := jwtx.ParseGoogleJWT(authReq.IdToken, claims, h.JWKS.Keyfunc)

	if parseErr != nil {
		log.Println(parseErr.Error())
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
	}

	verifyErr := jwtx.VerifyGoogleClaims(claims)
	if verifyErr != nil {
		log.Println(verifyErr.Error())
		httpx.WriteJSONError(w, "Unauthorized", http.StatusUnauthorized)
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
		log.Println(tokenErr.Error())
		httpx.WriteJSONError(w, "Internal Server Error", http.StatusInternalServerError)
	}

	authRes := authResponse{
		AccessToken: signedToken,
	}

	httpx.WriteJSONSuccess(w, authRes)
}
