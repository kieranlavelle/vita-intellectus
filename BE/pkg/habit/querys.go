package habit

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func getHabit(h *Habit, c *pgxpool.Pool) error {
	query := `
		SELECT
			h.user_id,
			h.name,
			h.days,
			h.tags,
			h.description,
			h.selected_stats,
			case
				when c.time_completed::date = now()::date then true
				else false
			end as completed
		FROM
			habits h FULL
			JOIN completed_habits c ON h.id = c.habit_id
		WHERE
			h.user_id=$1 AND h.id=$2
		ORDER BY c.time_completed DESC
		LIMIT 1
	`

	return c.QueryRow(context.Background(), query, h.UID, h.ID).Scan(
		&h.UID, &h.Name, &h.Days,
		&h.Tags, &h.Description, &h.SelectedStats,
		&h.Completed,
	)
}

func updateHabit(h *Habit, c *pgxpool.Pool) error {
	query := `
		UPDATE
			habits
		SET
			name=$1, tags=$2, description=$3, selected_stats=$4
		WHERE
			id=$5 AND user_id=$6
	`
	_, err := c.Exec(
		context.Background(), query, h.Name,
		h.Tags, h.Description, h.SelectedStats, h.ID, h.UID,
	)
	return err
}

func insertHabit(h *Habit, c *pgxpool.Pool) error {
	query := `
		INSERT INTO
			habits (user_id, name, tags, days, description, selected_stats)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING
			id
	`

	err := c.QueryRow(
		context.Background(), query, h.UID,
		h.Name, h.Tags, h.Days, h.Description, h.SelectedStats,
	).Scan(&h.ID)
	return err
}

func deleteHabit(h *Habit, c *pgxpool.Pool) error {
	query := `
		DELETE FROM
			habits
		WHERE
			id=$1 AND user_id=$2
	`
	_, err := c.Exec(context.Background(), query, h.ID, h.UID)
	return err
}

func completeHabit(h *Habit, notes string, c *pgxpool.Pool) error {
	query := `
		INSERT INTO
			completed_habits (habit_id, time_completed, notes)
		VALUES
			($1, $2, $3)
	`
	_, err := c.Exec(context.Background(), query, h.ID, time.Now().UTC(), notes)
	return err
}

func unCompleteHabit(h *Habit, c *pgxpool.Pool) error {
	query := `
		DELETE FROM
			completed_habits
		WHERE
			id IN (
					SELECT
						id
					FROM
						completed_habits
					WHERE
						habit_id=$1
					ORDER BY
						time_completed DESC
					LIMIT 1
			)
	`
	_, err := c.Exec(context.Background(), query, h.ID)
	return err
}

func habitCompletions(h *Habit, c *pgxpool.Pool) (pgx.Rows, error) {
	query := `
		SELECT
			habit_id, time_completed, notes
		FROM
			completed_habits
		WHERE
			habit_id=$1
		ORDER BY
			time_completed DESC
	`
	return c.Query(context.Background(), query, h.ID)
}

func userHabits(uid int, c *pgxpool.Pool) (pgx.Rows, error) {
	query := `
		SELECT
			DISTINCT ON (h.id)
			h.id,
			h.user_id,
			h.name,
			h.days,
			h.tags,
			h.description,
			h.selected_stats,
			case
				when c.time_completed :: date = now() :: date then true
				else false
			end as completed
		FROM
			habits h FULL
			JOIN completed_habits c ON h.id = c.habit_id
		WHERE
			h.user_id=$1
		ORDER BY h.id, c.time_completed DESC
	`

	return c.Query(context.Background(), query, uid)
}

func rolling28dayCompleted(hid int, c *pgxpool.Pool) (pgx.Rows, error) {
	query := `
		SELECT
			time_completed
		FROM
			completed_habits
		WHERE
			time_completed > current_date - interval '28' day
		AND
			habit_id=$1
	`

	return c.Query(context.Background(), query, hid)
}
