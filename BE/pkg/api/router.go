package api

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func methods(m ...string) (methods []string) {
	methods = append(methods, m...)
	if os.Getenv("MODE") == "DEV" {
		methods = append(methods, "OPTIONS")
	}
	return
}

// CreateRoutes forms a database connecting and sets up the API routes.
func CreateRoutes() {

	// form a connection to the database
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("error connecting to DB: %v\n", err)
	}
	defer pool.Close()

	// setup out logger
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	env := &Env{DB: pool}
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.Use(corsMiddleware)

	router.HandleFunc("/health", HealthCheck(env)).Methods(methods("GET")...)

	router.HandleFunc("/habit", AddHabit(env)).Methods(methods("POST")...)
	router.HandleFunc("/habit/{id:[0-9]+}", DeleteHabit(env)).Methods(methods("DELETE")...)
	router.HandleFunc("/habit/{id:[0-9]+}", GetHabit(env)).Methods(methods("GET")...)
	router.HandleFunc("/habit/{id:[0-9]+}", Update(env)).Methods(methods("PUT")...)
	router.HandleFunc("/habit/{id:[0-9]+}/complete", Complete(env)).Methods(methods("PUT")...)
	router.HandleFunc("/habit/{id:[0-9]+}/complete", UnComplete(env)).Methods(methods("DELETE")...)
	router.HandleFunc("/habit/{id:[0-9]+}/info", HabitInfo(env)).Methods(methods("GET")...)
	router.HandleFunc("/habit/{id:[0-9]+}/completions", Completions(env)).Methods(methods("GET")...)

	router.HandleFunc("/habits", Habits(env)).Methods(methods("GET")...)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8004", loggedRouter))
}
