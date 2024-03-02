package initializers

import "go-jwt/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
