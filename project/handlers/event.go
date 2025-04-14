package handlers

import (
	"first-app/db"
	"first-app/models"
)

func SaveEvent(e models.Event) {
	err := db.InsertEvent(e)
	if err != nil {
		panic(err)
	}
}

func GetAllEvents() []models.Event {
	events, err := db.GetAllEvents()
	if err != nil {
		panic(err)
	}

	return events
}
