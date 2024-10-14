package db

import (
	"database/sql"
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
