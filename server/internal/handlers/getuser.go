package handlers

import (
	"log"
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodbx"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
	"github.com/aDiThYa-808/persona-box/internal/jwtx"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	token, extractErr := httpx.ExtractToken(r)
	if extractErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
	}

	claims, validateErr := jwtx.ValidateAccessToken(token)
	if validateErr != nil {
		httpx.WriteJSONError(w, "unauthorized", http.StatusUnauthorized)
	}

	ctx := r.Context()

	userId := claims.Sub

	user, getErr := dynamodbx.GetUser(ctx, userId)

	if getErr != nil {
		log.Println(getErr)
		httpx.WriteJSONError(w, "couldnt find user", http.StatusNotFound)
	}

	httpx.WriteJSONSuccess(w, user)
}
