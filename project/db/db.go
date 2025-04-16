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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, name, port, timeZone)

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

//
//func createTables() {
//	createEventsTable := `
//	CREATE TABLE IF NOT EXISTS events (
//		id INTEGER PRIMARY KEY AUTOINCREMENT,
//		name TEXT NOT NULL,
//		description TEXT NOT NULL,
//		location TEXT NOT NULL,
//		dateTime DATE NOT NULL,
//		userId INTEGER
//	)
//	`
//
//	_, err := DB.Exec(createEventsTable)
//	if err != nil {
//		panic("Error creating events table" + err.Error())
//	}
//}
//
//func InsertEvent(e models.Event) error {
//	query := `
//		INSERT INTO events(name, description, location, dateTime, userId)
//		VALUES(?, ?, ?, ?, ?)
//	`
//	stmt, err := DB.Prepare(query)
//	if err != nil {
//		return err
//	}
//	defer func(stmt *sql.Stmt) {
//		err = stmt.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(stmt)
//	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
//	if err != nil {
//		return err
//	}
//
//	id, err := result.LastInsertId()
//	if err != nil {
//		return err
//	}
//	e.ID = id
//
//	return err
//}
//
//func GetAllEvents() ([]models.Event, error) {
//	query := "SELECT * FROM events"
//	rows, err := DB.Query(query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var events []models.Event
//	for rows.Next() {
//		event := models.Event{}
//		var dateTimeStr string
//
//		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserId)
//		if err != nil {
//			return nil, err
//		}
//		event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
//
//		events = append(events, event)
//	}
//
//	return events, nil
//}
