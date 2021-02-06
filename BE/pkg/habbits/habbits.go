package habbits

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/helpers"
	"github.com/kieranlavelle/vita-intellectus/pkg/persistence"
	"github.com/kieranlavelle/vita-intellectus/pkg/users"
)

// Habbit represents a habbit a user wants to set
type Habbit struct {
	HabbitID       int      `json:"habbit_id"`
	UserID         int      `json:"user_id"`
	Name           string   `json:"name"`
	Days           []string `json:"days"`
	CompletedToday bool     `json:"completed_today"`
}

// CompleteHabbitBody represents the body expected by the completeHabbit request
type CompleteHabbitBody struct {
	HabbitID int `json:"habbit_id"`
}

// createHabbit does validation and creates a habbit struct
func createHabbit(jsonStr []byte, user users.User) (Habbit, error) {

	habbit := Habbit{}

	requestBody := make(map[string]interface{})
	err := json.Unmarshal(jsonStr, &requestBody)
	if err != nil {
		return Habbit{}, err
	}

	if val, ok := requestBody["name"]; ok {
		habbit.Name = val.(string)
		habbit.UserID = user.UserID
	} else {
		return Habbit{}, errors.New("A habbit must have a name")
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

	err = conn.QueryRow(
		context.Background(),
		"insert into habbits (user_id, name) VALUES ($1, $2) RETURNING habbit_id",
		user.UserID, habbit.Name,
	).Scan(&habbit.HabbitID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal server error"})
	}

	habbit.UserID = user.UserID
	c.JSON(http.StatusOK, habbit)

}

// GetHabbits get's every habbit for a user
func GetHabbits(c *gin.Context) {

	var habbit Habbit
	var habbits []Habbit
	var lastCompleted sql.NullTime

	if conn, err := helpers.DatabaseConnection(c); err == nil {
		if user, err := helpers.RequestUser(c); err == nil {

			rows := persistence.HabbitsByUser(conn, user.UserID)
			for rows.Next() {
				err := rows.Scan(&habbit.HabbitID, &habbit.UserID, &habbit.Name, &lastCompleted)
				if err != nil {
					c.AbortWithStatus(http.StatusInternalServerError)
				}

				lYear, lMonth, lday := lastCompleted.Time.Date()
				year, month, day := time.Now().Date()

				validYear := lYear == year
				validMonth := lMonth == month
				validDay := lday == day

				if validDay && validMonth && validYear {
					habbit.CompletedToday = true
				} else {
					habbit.CompletedToday = false
				}

				// add the habbit into a slice
				habbits = append(habbits, habbit)
			}
		}
	}

	c.JSON(http.StatusOK, habbits)
}

// CompleteHabbit add's a habbit_completion
func CompleteHabbit(c *gin.Context) {
	if conn, err := helpers.DatabaseConnection(c); err == nil {
		completeHabbitBody := CompleteHabbitBody{}
		c.BindJSON(&completeHabbitBody)

		err := persistence.AddTrackedHabbit(conn, completeHabbitBody.HabbitID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		err = persistence.UpdateLastCompleted(conn, completeHabbitBody.HabbitID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}
