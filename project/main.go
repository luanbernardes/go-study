package main

import (
	"example.com/project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/events", getEvents)
	r.POST("/events", saveEvents)

	r.Run(":8099")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(200, events)
}

func saveEvents(c *gin.Context) {
	var event models.Event
	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event.ID = 1
	event.UserId = 123

	models.Event.Save(event)

	c.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
