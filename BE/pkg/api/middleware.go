package api

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		username := r.Header.Get("X-Authenticated-UserId")
		log.Printf("%v %v [X-Authenticated-UserID] %v\n", r.Method, r.RequestURI, username)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
