package httpx

import (
	"encoding/json"
	"net/http"
)

// returns http success response with provided content s
func WriteJSONSuccess(w http.ResponseWriter, s any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
