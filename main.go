package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEvnVariables()
	initializers.ConnectToGb()
	initializers.MigrateDB()
}

func main() {
	routes := gin.Default()
	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	routes.POST("/signup", controllers.SignUp)
	routes.POST("/signin", controllers.Signin)
	routes.GET("/validate", middleware.Middleware, controllers.Validate)
	routes.Run("localhost:8080")
}
