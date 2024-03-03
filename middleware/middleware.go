package middleware

import (
	"fmt"
	"go-jwt/initializers"
	"go-jwt/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Middleware(c *gin.Context) {
	/*
	 * Get the cookes off req
	 * Decode and validate the jwt
	 * check the exp of JWT
	 * Find the user with uuid fo token subject
	 * Attach the req
	 * forward the route to its destination (c.Next())
	 */

	JWTtokenString, err := c.Cookie("Authorization")

	if err != nil {
		// If cookies is not present
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Code snippet from doc
	token, err := jwt.Parse(JWTtokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// fmt.Println(claims["exp"], claims["sub"])
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// If JWT token has expired
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var user models.User
		initializers.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == uuid.Nil {
			// If user does not exists
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach the user with request
		c.Set("user", user)

		// Continue destination route
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
