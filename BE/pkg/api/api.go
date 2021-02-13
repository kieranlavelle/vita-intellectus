package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/habits"
	"github.com/kieranlavelle/vita-intellectus/pkg/users"
)

func connectToDatabase() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

// AddDatabaseConnection add's a connection to the request context
func AddDatabaseConnection(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConnection", conn)
		c.Next()
	}
}

// AddUser add's a user object to the request context
func AddUser(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Authenticated-UserId")
		user, err := users.GetUser(conn, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "Ok"})
	return
}

// CreateRoutes forms a database connecting and sets up the API routes.
func CreateRoutes() {

	//setup logging
	// log.SetOutput(gin.DefaultWriter)

	// form a connection to the database
	connection := connectToDatabase()
	defer connection.Close(context.Background())

	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s %s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.Request.Header.Get("X-Authenticated-UserId"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))

	router.Use(AddDatabaseConnection(connection))
	router.Use(AddUser(connection))
	router.Use(gin.Recovery())

	// version the api
	router.GET("/health_check", healthCheck)

	router.POST("/habit", habits.AddHabit)
	router.PUT("/habit/complete", habits.CompleteHabit)
	router.DELETE("/habit/:habitID", habits.DeleteHabit)
	router.PUT("/habit", habits.UpdateHabit)

	router.GET("/habits", habits.GetHabits)

	router.Run("0.0.0.0:8004")
}
