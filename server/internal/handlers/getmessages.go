package handlers

import (
	"log"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func GetSessionMessages(w http.ResponseWriter, r *http.Request) {
	token, extractErr := httpx.ExtractToken(r)

	if extractErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	_, validateErr := jwtx.ValidateAccessToken(token)
	if validateErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	sessionID := r.PathValue("id")
	if sessionID == "" {
		httpx.WriteJSONError(w, "session id not found", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	messages, getErr := dynamodbx.GetAllSessionMessages(ctx, sessionID)
	if getErr != nil {
		log.Println(getErr)
		httpx.WriteJSONError(w, "couldnt fetch messages", http.StatusInternalServerError)
		return
	}

	httpx.WriteJSONSuccess(w, messages)
}
