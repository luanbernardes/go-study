package db

import (
	"first-app/config"
	"first-app/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func InitDB() Dbinstance {
	host := config.GetEnv("DB_HOST")
	user := config.GetEnv("DB_USER")
	password := config.GetEnv("DB_PASSWORD")
	name := config.GetEnv("DB_NAME")
	port := config.GetEnv("DB_PORT")
	timeZone := config.GetEnv("TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host,
		user,
		password,
		name,
		port,
		timeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	// TODO remove before production
	db.AutoMigrate(&models.Event{})

	DB = Dbinstance{
		Db: db,
	}

	return DB
}
