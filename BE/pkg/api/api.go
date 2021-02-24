package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/habits"
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

	env := &habits.Env{DB: c}
	router := mux.NewRouter()

	// add the middleware we want
	router.Use(loggingMiddleware)

	router.HandleFunc("/habit", habits.AddHabit(env)).Methods("POST")
	router.HandleFunc("/habit/{id:[0-9]+}", habits.GetHabit(env)).Methods("GET")
	router.HandleFunc("/habit/{id:[0-9]+}", habits.CompleteHabit(env)).Methods("PUT")

	router.HandleFunc("/habits", habits.GetHabits(env)).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe("0.0.0.0:8004", nil)

	// router.POST("/habit", habits.AddHabit)
	// router.PUT("/habit/complete", habits.CompleteHabit)
	// router.DELETE("/habit/:habitID", habits.DeleteHabit)
	// router.PUT("/habit", habits.UpdateHabit)
	// router.GET("/habit/:habitID", habits.GetHabit)

	// router.GET("/habits", habits.GetHabits)

}
