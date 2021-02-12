package habbits

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/helpers"
)

// AddHabbit creates a new habbit for the user in the database
func AddHabbit(c *gin.Context) {

	var habbit Habbit

	if conn, err := helpers.DatabaseConnection(c); err == nil {
		if user, err := helpers.RequestUser(c); err == nil {

			err := c.BindJSON(&habbit)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
			}

			err = conn.QueryRow(
				context.Background(),
				"insert into habbits (user_id, name, days) VALUES ($1, $2, $3) RETURNING habbit_id",
				user.UserID, habbit.Name, habbit.Days,
			).Scan(&habbit.HabbitID)

			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			habbit.UserID = user.UserID
			habbit.setDueDates()
			c.JSON(http.StatusOK, habbit)
			return
		}
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

// GetHabbits get's every habbit for a user
func GetHabbits(c *gin.Context) {

	habbitsFilter, _ := c.GetQuery("filter")
	filter := habbitsFilter == "due"

	var habbit Habbit
	var lastCompleted sql.NullTime
	habbits := []Habbit{}

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("Failed to connect to DB: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	rows, err := HabbitsByUser(conn, user.UserID, filter)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			c.JSON(http.StatusOK, gin.H{
				"habbits": habbits,
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "failed to get user habbits.",
			})
			return
		}
	}

	for rows.Next() {
		err := rows.Scan(
			&habbit.HabbitID,
			&habbit.UserID,
			&habbit.Name,
			&habbit.Days,
			&lastCompleted,
		)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
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
		habbit = habbit.setDueDates()
		habbits = append(habbits, habbit)
	}

	c.JSON(http.StatusOK, gin.H{
		"habbits": habbits,
	})
}

// CompleteHabbit add's a habbit_completion
func CompleteHabbit(c *gin.Context) {
	if conn, err := helpers.DatabaseConnection(c); err == nil {
		completeHabbitBody := CompleteHabbitBody{}
		c.BindJSON(&completeHabbitBody)

		err := AddTrackedHabbit(conn, completeHabbitBody.HabbitID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		err = UpdateLastCompleted(conn, completeHabbitBody.HabbitID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}

// DeleteHabbit removes the habbit specified
func DeleteHabbit(c *gin.Context) {

	// form the DB connection
	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// get the user making the request
	user, err := helpers.RequestUser(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	habbitID, err := strconv.Atoi(c.Param("habbitID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"detail": "invalid habbit_id"},
		)
		return
	}

	//Delete habbit
	err = DBDeleteHabbit(conn, user.UserID, habbitID)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

// UpdateHabbit changes the variable properties of a habbit
func UpdateHabbit(c *gin.Context) {

	habbit := Habbit{}
	err := c.ShouldBindJSON(&habbit)
	if err != nil {
		log.Printf("Error parsing body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "expected valid habbit in body.",
		})
	}

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("Failed to connect to DB: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	err = DBUpdateHabbit(conn, user.UserID, habbit)
	if err != nil {
		switch err.(type) {
		case *HabbitNotFoundError:
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "please pass a valid habbit_id for a habbit you own.",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "error updating habbit. Please try again.",
			})
		}
		return
	}

	habbit = habbit.setDueDates()
	c.JSON(http.StatusOK, gin.H{
		"habbit": habbit,
	})

}
