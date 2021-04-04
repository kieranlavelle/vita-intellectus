package helpers

import (
	"errors"
	"net/http"
	"time"

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

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func DateEquals(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func DateInPast(old, new time.Time) bool {
	yOld, mOld, dOld := old.Date()
	yNew, mNew, dNew := new.Date()

	oldDate := time.Date(yOld, mOld, dOld, 0, 0, 0, 0, time.UTC)
	newDate := time.Date(yNew, mNew, dNew, 0, 0, 0, 0, time.UTC)
	return oldDate.Before(newDate)
}
