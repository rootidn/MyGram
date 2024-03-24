package database

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbPort   = "5432"
	dbname   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, dbPort)
	dns := config
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("success connecting to database")
	db.AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
