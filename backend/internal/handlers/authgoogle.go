package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/httpx"
)

type AuthRequest struct {
	IdToken string `json:"id_token"`
}

type AuthResponse struct {
	AUD            string `json:"aud"`
	ISS            string `json:"iss"`
	SUB            string `json:"sub"`
	Email          string `json:"email"`
	Email_Verified string `json:"email_verified"`
	Name           string `json:"name"`
	IAT            string `json:"iat"`
	EXP            string `json:"exp"`
	JTI            string `json:"jti"`
}

func GoogleAuthHandler(w http.ResponseWriter, r *http.Request) {
	var AuthReq AuthRequest
	decodeReqErr := json.NewDecoder(r.Body).Decode(&AuthReq)

	if decodeReqErr != nil || AuthReq.IdToken == "" {
		httpx.WriteJSONError(w, decodeReqErr.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	verifyUrl := "https://oauth2.googleapis.com/tokeninfo?id_token=" + AuthReq.IdToken

	resp, verifyErr := http.Get(verifyUrl)
	if verifyErr != nil {
		httpx.WriteJSONError(w, verifyErr.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != 200 {
		httpx.WriteJSONError(w, "Token is invalid, expired or malformed", http.StatusUnauthorized)
	}

	var AuthRes AuthResponse
	decodeResErr := json.NewDecoder(resp.Body).Decode(&AuthRes)
	if decodeResErr != nil {
		httpx.WriteJSONError(w, "Couldnt verify token", http.StatusInternalServerError)
		return
	}

	resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthRes)
}
