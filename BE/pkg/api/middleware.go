package api

import (
	"log"
	"net/http"
	"os"
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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if os.Getenv("MODE") == "DEV" {
			if r.Header.Get("Access-Control-Allow-Origin") == "" {
				w.Header().Add("Access-Control-Allow-Origin", "*")
			}
			if r.Header.Get("Access-Control-Allow-Headers") == "" {
				w.Header().Add("Access-Control-Allow-Headers", "*")
			}
			if r.Header.Get("Access-Control-Allow-Methods") == "" {
				w.Header().Add("Access-Control-Allow-Methods", "*")
			}

			if r.Method == "OPTIONS" {
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
