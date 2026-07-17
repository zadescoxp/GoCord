package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/zadescoxp/GoCord.git/internal/models"
)

func Connect(dbURL string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Ready 🚀")
	return db

}
