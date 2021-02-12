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

			// if there are no days set the default to all
			if len(habbit.Days) == 0 {
				habbit.Days = []string{"monday", "tueday", "wednesday", "thursday",
					"friday", "saturday", "sunday",
				}
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

	completedHabbits := []Habbit{}
	dueHabbits := []Habbit{}
	notDueHabbits := []Habbit{}

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
				"due":       dueHabbits,
				"completed": completedHabbits,
				"not_due":   notDueHabbits,
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

		// set the due dates
		habbit = habbit.setDueDates()

		if habbit.NextDue.NextDue == "Today" {
			if habbit.CompletedToday {
				completedHabbits = append(completedHabbits, habbit)
			} else {
				dueHabbits = append(dueHabbits, habbit)
			}
		} else {
			notDueHabbits = append(notDueHabbits, habbit)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"due":       dueHabbits,
		"completed": completedHabbits,
		"not_due":   notDueHabbits,
	})
}

// CompleteHabbit add's a habbit_completion
func CompleteHabbit(c *gin.Context) {

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("error connecting to database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
	}

	completeHabbitBody := CompleteHabbitBody{}
	err := c.BindJSON(&completeHabbitBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid request body.",
		})
	}

	err = AddTrackedHabbit(conn, completeHabbitBody.HabbitID)
	if err != nil {
		log.Printf("error adding tracked habbit to db: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
	}

	err = UpdateLastCompleted(conn, completeHabbitBody.HabbitID)
	if err != nil {
		log.Printf("error updating last completed: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
	}

	//TODO: Alter to have new Habbit returned
	c.JSON(http.StatusOK, gin.H{"detail": "success"})
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
