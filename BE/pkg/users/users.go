package users

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// User represents a user in the database
type User struct {
	UserID   int
	Username string
	Email    string
}

// GetUser get's the full user from the database using the username and email
func GetUser(conn *pgx.Conn, username string, email string) (User, error) {
	// need to insert the user into the database for this app
	var userID int
	err := conn.QueryRow(
		context.Background(),
		"insert into users (username, email) VALUES ($1, $2) RETURNING user_id",
		username, email,
	).Scan(&userID)

	if err != nil {
		return User{}, err
	}

	return User{UserID: userID, Username: username, Email: email}, nil
}
