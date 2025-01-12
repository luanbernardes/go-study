package handlers

import (
	"example.com/project/db"
	"example.com/project/models"
	"time"
)

var mockEvents = models.Event{
	ID:          1,
	Name:        "Event 3",
	Description: "Description 1",
	Location:    "Location 1",
	DateTime:    time.Now(),
	UserId:      1,
}

var eventsMock = []models.Event{mockEvents}

func SaveEvent(e models.Event) {
	err := db.InsertEvent(e)
	if err != nil {
		panic(err)
	}
}

func GetAllEvents() []models.Event {
	return eventsMock
}
