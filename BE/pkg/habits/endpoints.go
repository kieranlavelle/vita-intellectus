package habits

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/helpers"
)

// AddHabit creates a new habit for the user in the database
func AddHabit(c *gin.Context) {

	var habit Habit

	// validate our DB connection
	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("error getting DB connection: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("error getting user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	err = c.BindJSON(&habit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid request body.",
		})
		return
	}

	// if there are no days set the default to all
	if len(habit.Days) == 0 {
		habit.Days = []string{"monday", "tueday", "wednesday", "thursday",
			"friday", "saturday", "sunday",
		}
	}

	err = conn.QueryRow(
		context.Background(),
		"insert into habits (user_id, name, days) VALUES ($1, $2, $3) RETURNING id",
		user.ID, habit.Name, habit.Days,
	).Scan(&habit.ID)

	if err != nil {
		log.Printf("error when creating habit: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	habit.UserID = user.ID
	habit = habit.setDueDates()
	c.JSON(http.StatusOK, habit)
	return
}

// GetHabits get's every habbit for a user
func GetHabits(c *gin.Context) {

	completedHabits := []Habit{}
	dueHabits := []Habit{}
	notDueHabits := []Habit{}

	allHabits := []Habit{}

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("Failed to get DB connection: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	rows, err := DBHabitsByUser(conn, user.ID)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			c.JSON(http.StatusOK, gin.H{
				"due":       dueHabits,
				"completed": completedHabits,
				"not_due":   notDueHabits,
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "failed to get user habits.",
			})
			return
		}
	}

	for rows.Next() {

		habit := Habit{}

		err := rows.Scan(&habit.ID, &habit.UserID, &habit.Name, &habit.Days)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "please try again later.",
			})
			return
		}

		allHabits = append(allHabits, habit)
	}

	for _, habit := range allHabits {
		completedToday, err := DBCompletedHabitsToday(conn, habit.ID)
		if err != nil {
			log.Printf("error when getting habit count: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "please try again later",
			})
		}

		if completedToday > 0 {
			habit.Completed = true
		} else {
			habit.Completed = false
		}

		// set the due dates
		habit = habit.setDueDates()

		if habit.NextDue.NextDue == "Today" {
			if habit.Completed {
				completedHabits = append(completedHabits, habit)
			} else {
				dueHabits = append(dueHabits, habit)
			}
		} else {
			notDueHabits = append(notDueHabits, habit)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"due":       dueHabits,
		"completed": completedHabits,
		"not_due":   notDueHabits,
	})
}

// CompleteHabit add's a habbit_completion
func CompleteHabit(c *gin.Context) {

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("error connecting to database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	completeHabitBody := CompleteHabitBody{}
	err = c.BindJSON(&completeHabitBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid request body.",
		})
		return
	}

	err = AddTrackedHabit(conn, completeHabitBody.HabitID)
	if err != nil {
		log.Printf("error adding completed_habit to db: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	//TODO: Alter to have new Habbit returned
	c.JSON(http.StatusOK, gin.H{"detail": "success"})
}

// DeleteHabit removes the habbit specified
func DeleteHabit(c *gin.Context) {

	// form the DB connection
	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("failed to get DB connection: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	// get the user making the request
	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later.",
		})
		return
	}

	habitID, err := strconv.Atoi(c.Param("habitID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "invalid habit id",
		})
		return
	}

	//Delete habit
	err = DBDeleteHabit(conn, user.ID, habitID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": "please pass a habit_id for a habit you own",
		})
		return
	}

	c.String(http.StatusOK, "success")
}

// UpdateHabit changes the variable properties of a habbit
func UpdateHabit(c *gin.Context) {

	habit := Habit{}
	err := c.ShouldBindJSON(&habit)
	if err != nil {
		log.Printf("Error parsing body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "expected valid habit in body.",
		})
		return
	}

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("Failed to get DB connection: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	err = DBUpdateHabit(conn, user.ID, habit)
	if err != nil {
		switch err.(type) {
		case *HabitNotFoundError:
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "please pass a valid habit_id for a habit you own.",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "error updating habit. Please try again.",
			})
		}
		return
	}

	habit = habit.setDueDates()
	c.JSON(http.StatusOK, gin.H{
		"habit": habit,
	})
}

// HabitCompletions returns all completions for a given habbit
func HabitCompletions(c *gin.Context) {

	habitCompletions := HabitCompletionsResponse{}
	habitCompletionsRows := []HabitCompletedRow{}

	completions := HabitCompletionsBody{}
	err := c.ShouldBindJSON(&completions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid request body",
		})
		return
	}

	conn, err := helpers.DatabaseConnection(c)
	if err != nil {
		log.Printf("failed to get DB connection: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	user, err := helpers.RequestUser(c)
	if err != nil {
		log.Printf("failed to get user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "please try again later",
		})
		return
	}

	rows, err := dbCompletedHabits(conn, user.ID, completions.HabitID)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			c.JSON(http.StatusOK, habitCompletions)
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "please try again later",
			})
		}
		return
	}

	for rows.Next() {

		habitCompletion := HabitCompletedRow{}
		err := rows.Scan(&habitCompletion.HabitID, &habitCompletion.Time)
		if err != nil {
			log.Printf("failed to parse completed habit row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "please try again later",
			})
			return
		}

		habitCompletionsRows = append(habitCompletionsRows, habitCompletion)
	}

	c.JSON(http.StatusOK, habitCompletions)
}
