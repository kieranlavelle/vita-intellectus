package tasks

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func createEmptyTask(t Task, c *pgxpool.Pool) (Task, error) {
	query := `
		INSERT INTO
			tasks (uid, name, description, recurring, days, date, extra)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	err := c.QueryRow(context.Background(), query, t.UID, t.Name, t.Description,
		t.Recurring, t.Days, t.Date, t.Extra).Scan(&t.ID)

	return t, err
}
