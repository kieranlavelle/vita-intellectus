package habbits

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func getScheduleID(conn *pgx.Conn, schedule HabbitSchedule) (int, error) {

	scheduleID := -1
	err := conn.QueryRow(
		context.Background(),
		"SELECT schedule_id FROM schedules WHERE name=$1 AND user_id=$2",
		schedule.Name, schedule.UserID,
	).Scan(&scheduleID)

	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			break
		default:
			return -1, err
		}
	}

	// need to check for a no-rows error
	if err == pgx.ErrNoRows {
		err = conn.QueryRow(
			context.Background(),
			"INSERT INTO schedules (user_id, name, days, times) VALUES ($1, $2, $3, $4) RETURNING schedule_id",
			schedule.UserID, schedule.Name, schedule.Days, schedule.Times,
		).Scan(&scheduleID)

		if err != nil {
			return -1, err
		}
	}

	return scheduleID, nil
}

func getScheduleByID(conn *pgx.Conn, scheduleID *int) (HabbitSchedule, error) {

	schedule := HabbitSchedule{}
	err := conn.QueryRow(
		context.Background(),
		"SELECT * FROM schedules WHERE schedule_id=$1",
		scheduleID,
	).Scan(&schedule.ScheduleID, &schedule.UserID, &schedule.Name,
		&schedule.Days, &schedule.Times,
	)

	if err != nil {
		return HabbitSchedule{}, err
	}

	return schedule, nil
}
