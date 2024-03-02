package main

import (
	"go-jwt/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()
	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	routes.Run("localhost:8080")
}

func init() {
	initializers.LoadEvnVariables()
	initializers.ConnectToGb()
}
