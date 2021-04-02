package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/kieranlavelle/vita-intellectus/pkg/habit"
)

type CompleteBody struct {
	Notes string `json:"notes"`
}

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
			env.internalServerError(w, r, err)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", 500)
			return
		}

		habit := h.New(user.ID)
		err = json.Unmarshal(body, &habit)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "invalid request body",
			})
			return
		}

		err = habit.Update(env.DB)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(habit)
	}
}

// DeleteHabit removes a habit and it's completions from
// the database if it exists and is owned by the user. The
// checks for ownership and existence are done at load time.
func DeleteHabit(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		// ownership and existence checks are done prior to call
		// so any error is a server error
		err = habit.Delete(env.DB)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"detail": "deleted",
		})
	}
}

// GetHabit returns a habit from the database if it exists
// and is owned by the caller
func GetHabit(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := env.getUser(r)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]h.Habit{
			"habit": habit,
		})
	}
}

// Update updates the habit if the user owns it
// and has specified values that are updateable
func Update(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		err = json.Unmarshal(body, &habit)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "expected body as a valid habit",
			})
			return
		}

		// Perform some checks and if the specified updates are
		// valid update the database entry with the new values
		err = habit.Update(env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]h.Habit{
			"habit": habit,
		})
	}
}

// Complete updates the habit if the user owns it
// and add's a completion record for today
func Complete(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", 500)
			return
		}

		// don't need to check the unmarshal error as if
		// the body is incorrect notes will be "" which is
		// what we want
		completedNotes := CompleteBody{}
		json.Unmarshal(body, &completedNotes)

		// Perform some checks and if the specified updates are
		// valid update the database entry with the new values
		err = habit.Complete(completedNotes.Notes, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]h.Habit{
			"habit": habit,
		})
	}
}

// UnComplete updates the habit if the user owns it
// and removes a completion record for today
func UnComplete(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		// Perform some checks and if the specified updates are
		// valid update the database entry with the new values
		err = habit.UnComplete(env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]h.Habit{
			"habit": habit,
		})
	}
}

// Completions returns all of the times a habit
// has been completed
func Completions(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		// Perform some checks and if the specified updates are
		// valid update the database entry with the new values
		completions, err := habit.Completions(env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(completions)
	}
}

// Habits returns all of the habits for a user
func Habits(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := env.getUser(r)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		// Perform some checks and if the specified updates are
		// valid update the database entry with the new values
		habits, err := h.Habits(user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(map[string][]h.Habit{
			"habits": habits,
		})

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"detail": err.Error(),
			})
		}
	}
}

// HabitInfo returns all of the habits for a user
func HabitInfo(env *Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := env.getUser(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		// don't need to check the error as it is done
		// by the router.
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		habit, err := h.Load(id, user.ID, env.DB)
		if err != nil {
			switch err.(type) {
			case *h.Error:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{
					"detail": err.Error(),
				})
			default:
				env.internalServerError(w, r, err)
			}
			return
		}

		info, err := habit.Info(env.DB)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statistics": info,
		})
	}
}
