package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"ramache-zoubir/todo-list-api-golang/models"
)

var DbConnection *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(" error ", err)
	}
	createTables(db)
	DbConnection = db
}

func createTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Todo{}, &models.User{})
	if err != nil {
		log.Fatal("error when creating tables ")
	}
}
