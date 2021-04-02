package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	t "github.com/kieranlavelle/vita-intellectus/pkg/tasks"
)

// AddTask creates a new task for the user in the database
func AddTask(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := env.getUser(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail", "internal server error. Please try later.",
			})
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", 500)
			return
		}

		task := t.Task{UID: user.ID}
		err = json.Unmarshal(body, &task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"detail": "invalid request body",
			})
			return
		}

		task, err = t.New(task, env.DB)
		if err != nil {
			env.internalServerError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}
