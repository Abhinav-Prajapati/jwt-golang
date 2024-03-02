package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	Code  string
	Price uint
	Name  string
}

func ConnectToGb() {
	var err error
	dsn := os.Getenv("SUPABASE_PG_CONNECTION_STRING")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	} else {
		fmt.Println("Connected to DB")
	}
}
