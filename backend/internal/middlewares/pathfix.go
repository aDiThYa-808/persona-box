package middlewares

import (
	"log"
	"net/http"
	"strings"
)

func StripStagePrefix(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incomming path: %v", r.URL.Path)
		trimmedPath, exists := strings.CutPrefix(r.URL.Path, "/default")
		log.Printf("Trimmed prefix: %v", trimmedPath)
		if exists {
			r.URL.Path = trimmedPath
		}

		next.ServeHTTP(w, r)
	})
}
