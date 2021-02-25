package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// CreateRoutes forms a database connecting and sets up the API routes.
func CreateRoutes() {

	// form a connection to the database
	c, _ := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	err := c.Ping(context.Background())
	if err != nil {
		log.Fatal("failed to connect to the database.")
	}
	defer c.Close(context.Background())

	env := &Env{DB: c}
	router := mux.NewRouter()

	// add the middleware we want
	router.Use(loggingMiddleware)

	router.HandleFunc("/habit", AddHabit(env)).Methods("POST")
	router.HandleFunc("/habit/{id:[0-9]+}", DeleteHabit(env)).Methods("DELETE")
	router.HandleFunc("/habit/{id:[0-9]+}", GetHabit(env)).Methods("GET")
	router.HandleFunc("/habit/{id:[0-9]+}", Update(env)).Methods("PUT")
	router.HandleFunc("/habit/{id:[0-9]+}/complete", Complete(env)).Methods("PUT")
	router.HandleFunc("/habit/{id:[0-9]+}/completions", Completions(env)).Methods("GET")

	router.HandleFunc("/habits", Habits(env)).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe("0.0.0.0:8004", nil)
}
