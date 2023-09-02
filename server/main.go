package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("LOCAL_ENV") == "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to github workflow")
	})

	r.GET("/env", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":     os.Getenv("NAME"),
			"location": os.Getenv("LOCATION"),
		})
	})

	r.Run("0.0.0.0:8080")
}
