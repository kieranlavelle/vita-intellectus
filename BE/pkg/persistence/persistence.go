package persistence

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

// HabbitsByUser returns habbit_id,user_id,name
func HabbitsByUser(conn *pgx.Conn, userID int) pgx.Rows {
	rows, _ := conn.Query(
		context.Background(),
		"SELECT habbit_id, user_id, name, last_completed FROM habbits WHERE user_id=$1",
		userID,
	)
	return rows
}

// CompletedHabbitsToday returns the number of completed habbits against a habbit today
func CompletedHabbitsToday(conn *pgx.Conn, habbitID int) (int, error) {
	completedHabbits := 0
	err := conn.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM tracked_habbits WHERE habbit_id=$1 AND time_completed::date = now()::date;",
		habbitID,
	).Scan(&completedHabbits)

	if err != nil {
		return -1, err
	}

	return completedHabbits, nil
}

// AddTrackedHabbit Adds a row into the tracked_habbits table
func AddTrackedHabbit(conn *pgx.Conn, habbitID int) error {
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO tracked_habbit (habbit_id, time_completed) VALUES ($1, $2)",
		habbitID, time.Now().UTC(),
	)

	if err != nil {
		return err
	}

	return nil

}

// UpdateLastCompleted updates the last_completed field in the habbits table
func UpdateLastCompleted(conn *pgx.Conn, habbitID int) error {
	_, err := conn.Exec(
		context.Background(),
		"UPDATE habbits SET last_completed=$1 WHERE habbit_id=$2",
		time.Now().UTC(), habbitID,
	)

	if err != nil {
		return err
	}

	return nil
}
