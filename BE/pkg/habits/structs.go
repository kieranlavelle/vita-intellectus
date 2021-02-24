package habits

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/kieranlavelle/vita-intellectus/pkg/users"
)

type Env struct {
	DB *pgx.Conn
}

func (e *Env) getUser(r *http.Request) (users.User, error) {
	u := r.Header.Get("X-Authenticated-UserId")
	return users.GetUser(e.DB, u)
}

// Habit represents a habbit a user wants to set
type Habit struct {
	ID        int      `json:"id"`
	UserID    int      `json:"user_id"`
	Name      string   `json:"name"`
	Days      []string `json:"days"`
	Tags      []string `json:"tags"`
	NextDue   DueDates `json:"due_dates"`
	Completed bool     `json:"completed"`
}

// CompleteHabitBody represents the body expected by the completeHabbit request
type CompleteHabitBody struct {
	HabitID int `json:"habit_id"`
}

// DueDates represents when a habbit is next due depending on state
type DueDates struct {
	NextDue               string `json:"next_due"`
	NextDueAfterCompleted string `json:"next_due_on_completed"`
}

// HabitCompletionsBody represents the request body we expect
// to recieve when a request is made to fetch habbit completions
type HabitCompletionsBody struct {
	HabitID int `json:"habit_id"`
}

// HabitCompletedRow represents a row a the response when
// habbit completions are requested
type HabitCompletedRow struct {
	HabitID int       `json:"habit_id"`
	Time    time.Time `json:"time"`
}

// HabitCompletionsResponse represents the response body for
// the habit completions endpoint
type HabitCompletionsResponse struct {
	Completions []HabitCompletedRow `json:"completions"`
}
