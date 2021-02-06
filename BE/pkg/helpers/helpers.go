package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/users"
)

// ErrorExit sets the context and returns true when an error is found
func ErrorExit(err error, ctx *gin.Context, msg string, status int) bool {
	if err != nil {
		ctx.JSON(status, gin.H{"detail": msg})
		return true
	}
	return false
}

// DatabaseConnection gets the connection out of the context
func DatabaseConnection(c *gin.Context) (*pgx.Conn, error) {
	conn, ok := c.MustGet("databaseConnection").(*pgx.Conn)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
		return nil, errors.New("failed to get connection from context")
	}

	return conn, nil
}

// RequestUser gets the user out of the context
func RequestUser(c *gin.Context) (users.User, error) {
	user, ok := c.MustGet("user").(users.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
		return users.User{}, errors.New("failed to get user from context")
	}

	return user, nil
}
