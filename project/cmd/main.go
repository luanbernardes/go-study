package main

import (
	"net/http"

	"first-app/db"
	"first-app/handlers"
	"first-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.GET("/events", getEvents)
	r.POST("/events", saveEvents)

	err := r.Run(":8099")
	if err != nil {
		panic("Not could run in 8090 port" + err.Error())
	}
}

func getEvents(c *gin.Context) {
	events := handlers.GetAllEvents()
	c.JSON(200, events)
}

func saveEvents(c *gin.Context) {
	var event models.Event
	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	handlers.SaveEvent(event)

	c.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
