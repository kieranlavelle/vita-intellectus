package habits

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

func dbInsertNewHabit(db *pgx.Conn, h Habit, uid int) (int, error) {
	query := `INSERT INTO habits (user_id, name, days, tags)
		VALUES ($1, $2, $3, $4) RETURNING id
	`

	id := 0
	ctx := context.Background()
	err := db.QueryRow(ctx, query, uid, h.Name, h.Days, h.Tags).Scan(&id)

	return id, err
}

// DBHabitsByUser returns habbit_id,user_id,name
func DBHabitsByUser(conn *pgx.Conn, userID int) (pgx.Rows, error) {
	return conn.Query(
		context.Background(),
		"SELECT id, user_id, name, days, tags FROM habits WHERE user_id=$1",
		userID,
	)
}

// DBUpdateHabit updates a habbit in the DB
func DBUpdateHabit(conn *pgx.Conn, userID int, habit Habit) error {

	commandTag, err := conn.Exec(
		context.Background(),
		"UPDATE habits SET name=$1, days=$2 WHERE id=$3 AND user_id=$4",
		habit.Name, habit.Days, habit.ID, userID,
	)

	if commandTag.RowsAffected() == 0 {
		return &HabitNotFoundError{msg: "habit not found."}
	}
	return err
}

// DBCompletedHabitsToday returns the number of completed habbits against a habit today
func DBCompletedHabitsToday(conn *pgx.Conn, habitID int) (int, error) {
	query := `
		SELECT 
			COUNT(*)
		FROM
			completed_habits
		WHERE
			habit_id=$1 AND time_completed::date = now()::date
	`

	completedHabits := 0
	err := conn.QueryRow(context.Background(), query, habitID).Scan(&completedHabits)

	return completedHabits, err
}

// AddTrackedHabit Adds a row into the tracked_habbits table
func AddTrackedHabit(conn *pgx.Conn, habitID int) error {
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO completed_habits (habit_id, time_completed) VALUES ($1, $2)",
		habitID, time.Now().UTC(),
	)

	return err
}

// DBDeleteHabit removes the habit if the user owns in
func DBDeleteHabit(conn *pgx.Conn, userID int, habitID int) error {
	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM habits WHERE id=$1 AND user_id=$2",
		habitID, userID,
	)

	return err
}

func dbGetHabit(conn *pgx.Conn, userID int, habitID int) pgx.Row {
	query := `
		SELECT
			id, user_id, name, days, tags
		FROM
			habits
		WHERE
			user_id=$1 AND id=$2
	`

	return conn.QueryRow(context.Background(), query, userID, habitID)
}

func dbCompletedHabits(conn *pgx.Conn, userID int, habitID int) (pgx.Rows, error) {

	query := `
		SELECT
			*
		FROM 
			habits h
		FULL JOIN completed_habits c
			ON h.id=c.habbit_id
		WHERE
			h.user_id=$1 AND h.habit_id=$2`

	return conn.Query(
		context.Background(),
		query,
		userID, habitID,
	)

}

func dbInsertHabit(conn *pgx.Conn, userID int, name string, days, tags []string) (int, error) {
	query := `
		INSERT INTO
			habits
				(user_id, name, days, tags)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`

	id := 0
	err := conn.QueryRow(context.Background(), query, userID, name, days, tags).Scan(
		&id,
	)
	return id, err
}

func dbGetPastMonthCompletions(conn *pgx.Conn, userID, habitID int) (pgx.Rows, error) {
	query := `
		SELECT 
			habit_id, time_completed
		FROM
			completed_habits
		WHERE
			time_completed > current_date - interval '28' day
		AND
			habit_id=$1 AND user_id=$2
	`

	return conn.Query(context.Background(), query, userID, habitID)
}
