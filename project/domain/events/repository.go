package events

import "first-app/models"

type Repository interface {
	GetAllEvents() []models.Event
}
