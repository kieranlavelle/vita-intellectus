package api

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func TaskRouter() (*gin.Engine, *pgxpool.Pool) {

	// form a connection to the database
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("error connecting to DB: %v\n", err)
	}

	env := &Env{DB: pool}

	r := gin.Default()

	r.POST("/task", AddTask(env))
	r.GET("/task/:task_id", GetTask(env))
	r.GET("/tasks", GetTasks(env))
	r.PUT("/task/complete/:task_id", CompleteTask(env))

	return r, pool
}
