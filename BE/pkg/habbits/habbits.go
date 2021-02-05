package habbits

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
	"github.com/mitchellh/mapstructure"

	"github.com/kieranlavelle/vita-intellectus/pkg/helpers"
	"github.com/kieranlavelle/vita-intellectus/pkg/users"
)

// HabbitSchedule specifies times for a habbit to be formed
type HabbitSchedule struct {
	ScheduleID int      `json:"schedule_id"`
	UserID     int      `json:"user_id"`
	Name       string   `json:"name" validate:"required"`
	Days       []string `json:"days" validate:"required"`
	Times      []string `json:"times"`
}

// Habbit represents a habbit a user wants to set
type Habbit struct {
	HabbitID    int            `json:"habbit_id"`
	UserID      int            `json:"user_id"`
	Name        string         `json:"name"`
	HasSchedule bool           `json:"has_schedule"`
	Schedule    HabbitSchedule `json:"schedule"`
}

// createHabbit does validation and creates a habbit struct
func createHabbit(jsonStr []byte, user users.User) (Habbit, error) {

	validate := validator.New()

	requestBody := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &requestBody)
	if err != nil {
		return Habbit{}, err
	}

	habbit := Habbit{}
	schedule := HabbitSchedule{}

	if val, ok := requestBody["schedule"]; ok {
		mapstructure.Decode(val, &schedule)
		validate.Struct(&schedule)
		habbit.HasSchedule = true
	}

	if val, ok := requestBody["name"]; ok {
		habbit.Name = val.(string)
		habbit.UserID = user.UserID
		schedule.UserID = user.UserID
	} else {
		return Habbit{}, errors.New("A habbit must have a name")
	}

	if habbit.HasSchedule {
		habbit.Schedule = schedule
	}

	return habbit, nil

}

// AddHabbit creates a new habbit for the user in the database
func AddHabbit(c *gin.Context) {

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

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
	}

	habbit, err := createHabbit(jsonData, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
	}

	// Check if this schedule exists. If it does get ID
	if habbit.HasSchedule {
		scheduleID, err := getScheduleID(conn, habbit.Schedule)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		}
		habbit.Schedule.ScheduleID = scheduleID

		err = conn.QueryRow(
			context.Background(),
			"insert into habbits (user_id, name, schedule_id) VALUES ($1, $2, $3) RETURNING habbit_id",
			user.UserID, habbit.Name, habbit.Schedule.ScheduleID,
		).Scan(&habbit.HabbitID)

	} else {
		err = conn.QueryRow(
			context.Background(),
			"insert into habbits (user_id, name, schedule_id) VALUES ($1, $2) RETURNING habbit_id",
			user.UserID, habbit.Name,
		).Scan(&habbit.HabbitID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal server error"})
	}

	habbit.UserID = user.UserID

	c.JSON(http.StatusOK, habbit)

}

// GetHabbits get's every habbit for a user
func GetHabbits(c *gin.Context) {
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

	var habbits []Habbit
	var scheduleIDs []*int

	// need to load the habbitSchedule in
	rows, _ := conn.Query(
		context.Background(),
		"SELECT habbit_id, user_id, name, schedule_id FROM habbits WHERE user_id=$1",
		user.UserID,
	)

	for rows.Next() {
		habbit := Habbit{}
		var scheduleID *int
		err := rows.Scan(&habbit.HabbitID, &habbit.UserID, &habbit.Name, &scheduleID)
		if helpers.ErrorExit(err, c, "failed to get habbits", http.StatusBadRequest) {
			return
		}

		scheduleIDs = append(scheduleIDs, scheduleID)
		habbits = append(habbits, habbit)
	}

	for index, scheduleID := range scheduleIDs {
		if scheduleID != nil {
			schedule, err := getScheduleByID(conn, scheduleID)
			if helpers.ErrorExit(err, c, "failed to get habbit schedule", http.StatusBadRequest) {
				return
			}

			habbits[index].HasSchedule = true
			habbits[index].Schedule = schedule
		} else {
			habbits[index].HasSchedule = false
			habbits[index].Schedule = HabbitSchedule{}
		}
	}

	c.JSON(http.StatusOK, habbits)
}
