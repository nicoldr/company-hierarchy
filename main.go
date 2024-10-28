package main

import (
	"database/sql"
	"log"
	"os"

	"company-hierarchy/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Configure database connection
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()

    // Create a new Gin router
    router := gin.Default()

    // Define a base route
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Welcome to the Company Hierarchy API!"})
    })

    // Define routes for department management
    controllers.SetupRoutes(router, db)

    // Start the server on the specified port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if PORT is not set
    }
    router.Run(":" + os.Getenv("PORT"))
}
