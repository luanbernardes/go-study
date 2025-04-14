package db

import (
	"database/sql"
	"first-app/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "event.db")

	if err != nil {
		panic(err)
	}

	DB = db

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TEXT NOT NULL,
		userId INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Error creating events table" + err.Error())
	}
}

func InsertEvent(e models.Event) error {
	query := `
		INSERT INTO events(name, description, location, dateTime, userId) 
		VALUES(?, ?, ?, ?, ?)
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
}

func GetAllEvents() ([]models.Event, error) {
	query := "SELECT * FROM events"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		event := models.Event{}
		var dateTimeStr string

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTimeStr, &event.UserId)
		if err != nil {
			return nil, err
		}
		event.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)

		events = append(events, event)
	}

	return events, nil
}
