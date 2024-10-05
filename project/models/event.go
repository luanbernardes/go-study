package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var mockEvents = Event{
	ID:          1,
	Name:        "Event 3",
	Description: "Description 1",
	Location:    "Location 1",
	DateTime:    time.Now(),
	UserId:      1,
}

var events = []Event{mockEvents}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
