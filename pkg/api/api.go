package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func connectToDatabase() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	return conn
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
	router.Use(cors.Default())

	// version the api
	router.GET("/health_check", healthCheck)
	router.Run("0.0.0.0:8004")

}
