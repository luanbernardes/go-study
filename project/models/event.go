package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
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

func Save(e Event) {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
