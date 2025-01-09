package main

import (
	"net/http"

	"event-management/db"
	"event-management/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBind(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request data.",
		})
		return
	}

	event.Id = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created !",
		"event":   event,
	})
}
