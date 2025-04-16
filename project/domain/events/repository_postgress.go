package events

import (
	"first-app/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetAllEvents() []models.Event {
	var events []models.Event
	r.db.Find(&events)
	return events
}
