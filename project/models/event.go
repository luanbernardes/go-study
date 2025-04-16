package models

import (
	"time"
)

type Event struct {
	ID          int64     `gorm:"primarykey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId"`
}
