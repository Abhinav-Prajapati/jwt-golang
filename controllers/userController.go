package controllers

import (
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Note : hand user creation with same email ( tho email is set to be unique in db)
	// get the email and password
	var body struct {
		Email    string
		Password string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
	}
	// create user
	user_uuid := uuid.New()
	user := models.User{ID: user_uuid, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
	}

	c.JSON(http.StatusOK, gin.H{"id": user_uuid})

}

func Signin(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	// get request body
	c.Bind(&body)

	var user models.User
	// find user and store it row data to user struct
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email id ",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"user": "incorrect password"})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": "correct password"})
	}
}
