// database/db.go
package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"main/models"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	DB = db

	// Auto-migrate models
	err = DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	log.Println("Database connection established and models migrated successfully")
}
