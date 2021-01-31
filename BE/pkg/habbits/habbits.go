package habbits

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/users"
)

// HabbitSchedule specifies times for a habbit to be formed
type HabbitSchedule struct {
	ScheduleID int      `json:"schedule_id"`
	Name       string   `json:"name"`
	Days       []string `json:"days"`
	Times      []string `json:"times"`
}

// Habbit represents a habbit a user wants to set
type Habbit struct {
	HabbitID int            `json:"habbit_id"`
	UserID   int            `json:"user_id"`
	Name     string         `json:"name"`
	Schedule HabbitSchedule `json:"schedule"`
}

// CreateHabbit creates a new habbit for the user in the database
func CreateHabbit(c *gin.Context) {

	conn, ok := c.MustGet("databaseConnection").(*pgx.Conn)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
		return
	}
	user, ok := c.MustGet("user").(users.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
		return
	}

	habbit := Habbit{}
	err := c.ShouldBindJSON(&habbit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid request body."})
		return
	}

	habbit.UserID = user.UserID
	err = conn.QueryRow(
		context.Background(),
		"insert into habbits (user_id, name, schedule_id) VALUES ($1, $2) RETURNING habbit_id",
		user.UserID, habbit.Name, -1,
	).Scan(&habbit.HabbitID)

}
