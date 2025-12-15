package handlers

import (
	"net/http"

	"github.com/aDiThYa-808/persona-box/internal/dynamodb"
	"github.com/aDiThYa-808/persona-box/internal/httpx"
)

func Tstendpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := dynamodb.DescribeTable(ctx, "User")

	if err != nil {
		httpx.WriteJSONError(w, err.Error(), http.StatusInternalServerError)
	}

	httpx.WriteJSONSuccess(w, map[string]string{
		"message": "success",
	})
}
