package habits

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

// HealthCheck returns a 200 status
func HealthCheck(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}

// AddHabit creates a new habit for the user in the database
func AddHabit(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", 500)
			return
		}

		habit := Habit{Days: allDays, Tags: []string{}, UserID: user.ID}
		err = json.Unmarshal(body, &habit)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "invalid request body",
			})
			return
		}

		err = dbInsertNewHabit(env.DB, &habit, user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "please try again later",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(habit)
	}
}

// GetHabits returns all of a users habits
func GetHabits(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		rows, err := dbGetUserHabits(env.DB, user.ID)
		if err != nil {
			switch err {
			case pgx.ErrNoRows:
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string][]string{
					"habits": {},
				})
			default:
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": "internal server error, please try later",
				})
			}
			return
		}

		h := Habit{}
		habits := []Habit{}

		defer rows.Close()
		for rows.Next() {

			err := rows.Scan(&h.ID, &h.UserID, &h.Name,
				&h.Days, &h.Tags, &h.Completed,
			)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": "internal server error, please try later",
				})
				return
			}
			habits = append(habits, h)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string][]Habit{
			"habits": habits,
		})
	}
}

// GetHabit get's a single habit
func GetHabit(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		h := Habit{ID: id}
		err = dbGetHabit(env.DB, user.ID, &h)

		if err != nil {
			switch err {
			case pgx.ErrNoRows:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
			default:
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"detail": "internal server error. Please try later.",
				})
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]Habit{
			"habit": h,
		})
	}
}

// CompleteHabit add's a record to habit_completion table
func CompleteHabit(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		owned, err := dbCheckHabitIsOwned(env.DB, user.ID, id)
		if owned != true {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "please specify a habit you own",
			})
			return
		}

		err = dbCompleteHabit(env.DB, id)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "internal server error",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"detail": "completed",
		})
	}
}

// DeleteHabit removes the habbit specified
func DeleteHabit(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		owned, err := dbCheckHabitIsOwned(env.DB, user.ID, id)
		if owned != true {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "please specify a habit you own",
			})
			return
		}

		err = dbDeleteHabit(env.DB, user.ID, id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "internal server error",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"detail": "deleted",
		})

	}
}

// HabitCompletions returns all completions for a given habbit
func HabitCompletions(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		owned, err := dbCheckHabitIsOwned(env.DB, user.ID, id)
		if owned != true {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "please specify a habit you own",
			})
			return
		}

		rows, err := dbGetHabitCompletions(env.DB, id)
		if err != nil {
			switch err {
			case pgx.ErrNoRows:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string][]string{
					"completions": {},
				})
			}
		}

		for rows.Next() {
			...
		}

	}
}

// // HabitCompletions returns all completions for a given habbit
// func HabitCompletions(c *gin.Context) {

// 	habitCompletions := HabitCompletionsResponse{}
// 	habitCompletionsRows := []HabitCompletedRow{}

// 	completions := HabitCompletionsBody{}
// 	err := c.ShouldBindJSON(&completions)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"detail": "Invalid request body",
// 		})
// 		return
// 	}

// 	conn, err := helpers.DatabaseConnection(c)
// 	if err != nil {
// 		log.Printf("failed to get DB connection: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"detail": "please try again later",
// 		})
// 		return
// 	}

// 	user, err := helpers.RequestUser(c)
// 	if err != nil {
// 		log.Printf("failed to get user: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"detail": "please try again later",
// 		})
// 		return
// 	}

// 	rows, err := dbCompletedHabits(conn, user.ID, completions.HabitID)
// 	if err != nil {
// 		switch err {
// 		case pgx.ErrNoRows:
// 			c.JSON(http.StatusOK, habitCompletions)
// 		default:
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"detail": "please try again later",
// 			})
// 		}
// 		return
// 	}

// 	for rows.Next() {

// 		habitCompletion := HabitCompletedRow{}
// 		err := rows.Scan(&habitCompletion.HabitID, &habitCompletion.Time)
// 		if err != nil {
// 			log.Printf("failed to parse completed habit row: %v\n", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"detail": "please try again later",
// 			})
// 			return
// 		}

// 		habitCompletionsRows = append(habitCompletionsRows, habitCompletion)
// 	}

// 	c.JSON(http.StatusOK, habitCompletions)
// }

// // HabitInfo returns statistics about a given habit
// func HabitInfo(c *gin.Context) {
// 	conn, err := helpers.DatabaseConnection(c)
// 	if err != nil {
// 		log.Printf("Failed to get DB connection: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"detail": "please try again later.",
// 		})
// 		return
// 	}

// 	user, err := helpers.RequestUser(c)
// 	if err != nil {
// 		log.Printf("Failed to get user: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"detail": "please try again later.",
// 		})
// 		return
// 	}

// 	habitID, err := strconv.Atoi(c.Param("habitID"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"detail": "invalid habit id",
// 		})
// 		return
// 	}

// 	rows, err := dbGetPastMonthCompletions(conn, user.ID, habitID)
// 	if err != nil {
// 		switch err {
// 		case pgx.ErrNoRows:
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"detail": "please specify a habit you own",
// 			})
// 			return
// 		default:
// 			log.Printf("error when getting habit completions: %v\n", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"detail": "please try again later.",
// 			})
// 			return
// 		}
// 	}

// 	for rows.Next() {
// 		//
// 	}
// }
