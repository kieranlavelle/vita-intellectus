package tasks

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func createEmptyTask(t *Task, c *pgxpool.Pool) (*Task, error) {
	query := `
		INSERT INTO
			tasks (uid, name, description, recurring, days, date, date_created, extra)
		VALUES
			($1, $2, $3, $4, $5, $6, current_timestamp, $8)
		RETURNING id;
	`

	err := c.QueryRow(context.Background(), query, t.UID, t.Name,
		t.Description, t.Recurring, t.Days, t.Date, t.Extra).Scan(&t.ID)

	return t, err
}

func getTask(id, uid int, c *pgxpool.Pool) (*Task, error) {
	query := `
		SELECT
			name, description, recurring, days, date, date_created, extra
		FROM
			tasks
		WHERE
			id=$1 AND uid=$2
	`

	t := &Task{ID: id, UID: uid}
	err := c.QueryRow(context.Background(), query, id, uid).Scan(
		&t.Name, &t.Description, &t.Recurring, &t.Days,
		&t.Date, &t.DateCreated, &t.Extra,
	)

	return t, err
}

func getTasks(uid int, c *pgxpool.Pool) ([]*Task, error) {
	query := `
		SELECT
			id, name, description, recurring, days, date, date_created, extra
		FROM
			tasks
		WHERE
			uid=$1
	`

	tasks := make([]*Task, 0)
	rows, err := c.Query(context.Background(), query, uid)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()

	for rows.Next() {
		t := &Task{UID: uid}
		err = rows.Scan(&t.ID, &t.Name, &t.Description,
			&t.Recurring, &t.Days, &t.Date, &t.DateCreated, &t.Extra,
		)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, t)
	}

	return tasks, err
}

func completeTask(t *Task, notes string, date time.Time, c *pgxpool.Pool) error {
	query := `
		INSERT INTO
			task_completions (tid, notes, date)
		VALUES
			($1, $2, $3)
	`

	_, err := c.Exec(context.Background(), query, t.ID, notes, date)
	return err
}

func unCompleteTask(t *Task, date time.Time, c *pgxpool.Pool) error {
	query := `
		DELETE FROM
			task_completions
		WHERE
			tid=$1
		AND
			date_part('year', date)=$2
		AND
			date_part('month', date)=$3
		AND
			date_part('day', date)=$4
	`

	y, mString, d := date.Date()
	m := int(mString)

	_, err := c.Exec(context.Background(), query, t.ID, y, m, d)
	return err
}

func checkAbsoluteCompletion(id, d, m, y int, c *pgxpool.Pool) (bool, error) {
	query := `
		SELECT
			COUNT(*) > 0 as exists
		FROM
			task_completions
		WHERE
			tid=$1
		AND
			date_part('year', date)=$2
		AND
			date_part('month', date)=$3
		AND
			date_part('day', date)=$4
	`

	completed := false
	err := c.QueryRow(context.Background(), query, id, y, m, d).Scan(&completed)
	return completed, err
}

func updateTask(t *Task, c *pgxpool.Pool) error {
	query := `
		UPDATE
			tasks
		SET
			name=$1, description=$2
		WHERE
			id=$3
		AND
			uid=$4
	`

	_, err := c.Exec(context.Background(), query, t.Name, t.Description,
		t.ID, t.UID,
	)
	return err
}

func deleteTask(t *Task, c *pgxpool.Pool) error {
	query := `
		DELETE FROM
			tasks
		WHERE
			id=$1
		AND
			uid=$2
	`

	_, err := c.Exec(context.Background(), query, t.ID, t.UID)
	return err
}

func getNumCompletions(t *Task, date time.Time, c *pgxpool.Pool) (int, error) {
	query := `
		SELECT COUNT(*) FROM
			tasks
		WHERE
			id=$1
		AND
			uid=$2
		AND
			date<$3
	`

	count := 0
	err := c.QueryRow(context.Background(), query, t.ID, t.UID, date).Scan(
		&count,
	)
	return count, err
}
