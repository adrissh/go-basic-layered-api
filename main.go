package main

import (
	"GoLayeredCRUD/routes"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	// Muat file .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	// handle routes not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     "error",
			"statusCode": http.StatusNotFound,
			"message":    "Route not found",
			"path":       c.Request.URL.Path,
			"timestamp":  time.Now().Format(time.RFC3339),
		})
	})
	routes.EmployeeRoutes(r)

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}

}
