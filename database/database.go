package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("people.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	DATABASE = database
	log.Println("Database connected successfully")
}
