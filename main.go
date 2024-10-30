package main

import (
	"log"
	"os"

	"company-hierarchy/controllers"
	"company-hierarchy/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to the database
    database.Connect()
    defer database.DB.Close()

    // Create a new Gin router
    router := gin.Default()

    // Login endpoint
    router.POST("/login", controllers.Login)

    // JWT Authentication
    authorized := router.Group("/")
    authorized.Use(controllers.AuthenticateJWT())

    // Define a base route
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Welcome to the Company Hierarchy API!"})
    })

    // Define routes for department management
    controllers.SetupRoutes(authorized, database.DB)  // Protected routes (JWT)

    // Start the server on the specified port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if PORT is not set
    }
    router.Run(":" + port)
}
