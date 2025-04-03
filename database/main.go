package database

import (
	"axo/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Gorm Document
https://gorm.io/docs/
*/

var DB *gorm.DB

func Init() {
	// Load environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"))

	var err error

	// Connect to the database
	fmt.Println("🔌 Connecting to the database...")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	}
	fmt.Println("✅ Connected to the database")

	// Gorm Auto Migration operation
	// This will create the tables in the database if they do not exist.
	DB.AutoMigrate(
		//Demo Note application
		&models.Note{},

		// ⚠️ Axo Rest API Schemas ⚠️
		// 🎭 Auth & Role System
		&models.User{},
		&models.Role{},
		&models.Permission{},
	)
}
