package users

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// User represents a user in the database
type User struct {
	ID       int
	Username string
}

// GetUser get's the full user from the database using the username and email
func GetUser(conn *pgx.Conn, username string) (User, error) {
	// need to insert the user into the database for this app
	userID := -1
	err := conn.QueryRow(
		context.Background(),
		"SELECT user_id FROM users where username=$1",
		username,
	).Scan(&userID)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			break
		default:
			return User{}, err
		}
	}

	if userID != -1 {
		return User{ID: userID, Username: username}, nil
	}

	err = conn.QueryRow(
		context.Background(),
		"insert into users (username) VALUES ($1) RETURNING user_id",
		username,
	).Scan(&userID)

	if err != nil {
		return User{}, err
	}

	return User{ID: userID, Username: username}, nil
}
