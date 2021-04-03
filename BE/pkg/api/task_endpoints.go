package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"

	t "github.com/kieranlavelle/vita-intellectus/pkg/tasks"
)

// AddTask creates a new task for the user in the database
func AddTask(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := env.getUser(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error. Please try later.",
			})
			return
		}

		task := &t.Task{UID: user.ID}
		err = c.ShouldBindJSON(task)
		if err != nil {
			logrus.Errorf("error creating task: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "bad request body.",
			})
			return
		}

		task, err = t.New(task, env.DB)
		if err != nil {
			logrus.Errorf("error creating task in database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "failed to create task, please try later.",
			})
			return
		}

		task, err = t.Load(task.ID, user.ID, time.Now(), env.DB)
		if err != nil {
			logrus.Errorf("error loading task from database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "failed to create task, please try later.",
			})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

// GetTask returns a users task from the DB
func GetTask(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := env.getUser(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error. Please try later.",
			})
			return
		}

		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "please provide a valid task_id.",
			})
			return
		}

		dateString := c.Query("date")
		date := time.Now()
		if dateString != "" {
			layout := "2006-01-02T15:04:05.000Z"
			date, err = time.Parse(layout, dateString)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"detail": "please provide a valid date.",
				})
				return
			}
		}

		task, err := t.Load(task_id, user.ID, date, env.DB)
		if err != nil {
			logrus.Errorf("error loading task: %v", err)
			switch err {
			case pgx.ErrNoRows:
				c.JSON(http.StatusNotFound, gin.H{
					"detail": "not found.",
				})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"detail": "failed to get task. please try later",
				})
				return
			}
		}

		c.JSON(http.StatusOK, task)
	}
}

// GetTasks returns a users tasks from the DB
func GetTasks(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := env.getUser(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error. Please try later.",
			})
			return
		}

		dateString := c.Query("date")
		date := time.Now()
		if dateString != "" {
			layout := "2006-01-02T15:04:05.000Z"
			date, err = time.Parse(layout, dateString)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"detail": "please provide a valid date.",
				})
				return
			}
		}

		tasks, err := t.Tasks(user.ID, date, env.DB)
		if err != nil {
			logrus.Errorf("error loading task: %v", err)
			switch err {
			case pgx.ErrNoRows:
				c.JSON(http.StatusNotFound, gin.H{
					"detail": "not found.",
				})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"detail": "failed to get task. please try later",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	}
}

// CompleteTask returns a users task from the DB
func CompleteTask(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := env.getUser(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error. Please try later.",
			})
			return
		}

		// if this doesn't bind then no notes are provided
		completion := &TaskCompletion{}
		c.ShouldBindJSON(completion)

		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "please provide a valid task_id.",
			})
			return
		}

		dateString := c.Query("date")
		date := time.Now()
		if dateString != "" {
			layout := "2006-01-02T15:04:05.000Z"
			date, err = time.Parse(layout, dateString)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"detail": "please provide a valid date.",
				})
				return
			}
		}

		task, err := t.Load(task_id, user.ID, date, env.DB)
		if err != nil {
			logrus.Errorf("error loading task: %v", err)
			switch err {
			case pgx.ErrNoRows:
				c.JSON(http.StatusNotFound, gin.H{
					"detail": "not found.",
				})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"detail": "failed to get task. please try later",
				})
				return
			}
		}

		task, err = task.Complete(completion.Notes, date, env.DB)
		if err != nil {
			switch err.(type) {
			case *t.DisplayableError:
				c.JSON(http.StatusBadRequest, gin.H{
					"detail": err.Error(),
				})
				return
			default:
				logrus.Errorf("failed to complete task: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"detail": "failed to complete task.",
				})
				return
			}
		}

		c.JSON(http.StatusOK, task)
	}
}

// UnCompleteTask returns a users task from the DB
func UnCompleteTask(env *Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := env.getUser(c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"detail": "internal server error. Please try later.",
			})
			return
		}

		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"detail": "please provide a valid task_id.",
			})
			return
		}

		dateString := c.Query("date")
		date := time.Now()
		if dateString != "" {
			layout := "2006-01-02T15:04:05.000Z"
			date, err = time.Parse(layout, dateString)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"detail": "please provide a valid date.",
				})
				return
			}
		}

		task, err := t.Load(task_id, user.ID, date, env.DB)
		if err != nil {
			logrus.Errorf("error loading task: %v", err)
			switch err {
			case pgx.ErrNoRows:
				c.JSON(http.StatusNotFound, gin.H{
					"detail": "not found.",
				})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"detail": "failed to get task. please try later",
				})
				return
			}
		}

		task, err = task.UnComplete(date, env.DB)
		if err != nil {
			switch err.(type) {
			case *t.DisplayableError:
				c.JSON(http.StatusBadRequest, gin.H{
					"detail": err.Error(),
				})
				return
			default:
				logrus.Errorf("failed to complete task: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"detail": "failed to complete task.",
				})
				return
			}
		}

		c.JSON(http.StatusOK, task)
	}
}
