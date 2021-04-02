package tasks

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Tast represents a task a user owns
type Task struct {
	ID          int                    `json:"id"`
	UID         int                    `json:"user_id"`
	Name        int                    `json:"name"`
	Description string                 `json:"description"`
	Recurring   bool                   `json:"recurring"`
	Days        []string               `json:"days,omitempty"`
	Date        time.Time              `json:"date,omitempty"`
	Extra       map[string]interface{} `json:"extra"`
}

func New(t Task, c *pgxpool.Pool) (Task, error) {
	return createEmptyTask(t, c)
}

// func Load(id, uid int, c *pgxpool.Pool) (Task, error) {

// 	t := Task{}
// }
