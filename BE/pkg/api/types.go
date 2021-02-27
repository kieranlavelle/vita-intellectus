package api

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kieranlavelle/vita-intellectus/pkg/users"
	log "github.com/sirupsen/logrus"
)

// Env represents the context that is passed into
// each endpoint as a closure
type Env struct {
	DB *pgxpool.Pool
}

func (e *Env) getUser(r *http.Request) (users.User, error) {
	u := r.Header.Get("X-Authenticated-UserId")
	return users.GetUser(e.DB, u)
}

func (e *Env) internalServerError(w http.ResponseWriter, err error) {

	log.Fatalf("error connecting to DB: %v\n", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"detail": "internal server error, please try again later",
	})
}
