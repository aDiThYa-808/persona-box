package handlers

import (
	"log"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func GetChatSessions(w http.ResponseWriter, r *http.Request) {
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

	personaid := r.URL.Query().Get("personaid")
	if personaid == "" {
		httpx.WriteJSONError(w, "missing personaid", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	chats, getErr := dynamodbx.GetAllChatSessionsOfPersona(ctx, personaid)
	if getErr != nil {
		log.Println(getErr)
		httpx.WriteJSONError(w, "couldnt find chat sessions", http.StatusNotFound)
		return
	}

	httpx.WriteJSONSuccess(w, chats)

}
