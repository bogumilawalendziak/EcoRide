package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitDB() *gorm.DB {

	godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", host, user, dbname, password)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	err = db.AutoMigrate(&Vehicle{}, &User{}, &Reservation{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
