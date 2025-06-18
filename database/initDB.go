package database

import (
	"DevConnector/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	connStr := os.Getenv("DB")
	var err error

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.Profile{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}
