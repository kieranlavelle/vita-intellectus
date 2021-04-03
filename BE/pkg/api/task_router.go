package api

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func TaskRouter() (*gin.Engine, *pgxpool.Pool) {

	// form a connection to the database
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("error connecting to DB: %v\n", err)
	}

	env := &Env{DB: pool}

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/task", AddTask(env))
	r.GET("/task/:task_id", GetTask(env))
	r.DELETE("/task/:task_id", DeleteTask(env))
	r.GET("/tasks", GetTasks(env))
	r.PUT("/task/edit/:task_id", EditTask(env))
	r.PUT("/task/complete/:task_id", CompleteTask(env))
	r.PUT("/task/uncomplete/:task_id", UnCompleteTask(env))

	return r, pool
}
