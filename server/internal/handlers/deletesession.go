package handlers

import (
	"log"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func DeleteChatSession(w http.ResponseWriter, r *http.Request) {
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
	personaid := r.PathValue("pid")
	sessionid := r.PathValue("sid")
	if sessionid == "" {
		httpx.WriteJSONError(w, "missing sessionid", http.StatusBadRequest)
		return
	}

	deleteErr := dynamodbx.DeleteSession(r.Context(), personaid, sessionid)
	if deleteErr != nil {
		log.Println(deleteErr)
		httpx.WriteJSONError(w, "couldnt complete persona deletion", http.StatusInternalServerError)
		return
	}

	httpx.WriteJSONSuccess(w, "successfully deleted chat session")
}
