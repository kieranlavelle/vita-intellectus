package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/kieranlavelle/vita-intellectus/pkg/habbits"
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
		if c.Request.Method != "OPTIONS" {
			c.Set("databaseConnection", conn)
		}
		c.Next()
	}
}

// AddUser add's a user object to the request context
func AddUser(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			username := c.GetHeader("X-Authenticated-UserId")
			user, err := users.GetUser(conn, username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error."})
				return
			}
			c.Set("user", user)
		}
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
	log.SetOutput(gin.DefaultWriter)

	// form a connection to the database
	connection := connectToDatabase()
	defer connection.Close(context.Background())

	router := gin.Default()
	router.Use(AddDatabaseConnection(connection))
	router.Use(AddUser(connection))

	// version the api
	router.GET("/health_check", healthCheck)

	router.POST("/habbits", habbits.AddHabbit)
	router.GET("/habbits", habbits.GetHabbits)
	router.PUT("/habbit", habbits.UpdateHabbit)

	//depricate this
	router.PUT("/habbits/complete", habbits.CompleteHabbit)

	router.DELETE("/habbit/:habbitID", habbits.DeleteHabbit)

	router.Run("0.0.0.0:8004")
}
