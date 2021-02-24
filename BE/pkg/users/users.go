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
func GetUser(db *pgx.Conn, username string) (User, error) {
	// need to insert the user into the database for this app

	ctx := context.Background()
	query := "SELECT user_id FROM users where username=$1"
	user := User{}
	err := db.QueryRow(ctx, query, username).Scan(&user.ID)

	if err != nil {
		query = "insert into users (username) VALUES ($1) RETURNING user_id"
		err = db.QueryRow(ctx, query, username).Scan(&user.ID)
	}

	return user, err
}
