package initializers

import (
	"BookStore/models"
	"context"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Q models.Query
var CTX context.Context

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting to DB")
	}
	Q = *models.Use(DB)
}

func ConnectToLoggerDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if err := InitLogger(mongoURI, "bookstore_logs", "logs"); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	} else {
		log.Println("Logger initialized successfully")
	}
}
