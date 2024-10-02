package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/events", getEvents)
	r.Run(":8099")
}

func getEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "1",
	})
}
